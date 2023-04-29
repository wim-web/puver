package cmd

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wim-web/puver/internal/puver"
)

var (
	token         string
	org           string
	providerName  string
	pubKeyPath    string
	shasumPath    string
	shasumSigPath string
	version       string
	osName        string
	arch          string
	binaryPath    string
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy",
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeployToTerraformCloud(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&token, "token", "t", "", "")
	deployCmd.Flags().StringVarP(&org, "org", "o", "", "")
	deployCmd.Flags().StringVarP(&providerName, "name", "n", "", "")
	deployCmd.Flags().StringVar(&pubKeyPath, "pubkey-path", "", "")
	deployCmd.Flags().StringVar(&shasumPath, "shasum-path", "", "")
	deployCmd.Flags().StringVar(&shasumSigPath, "shasum-sig-path", "", "")
	deployCmd.Flags().StringVarP(&version, "version", "v", "", "")
	deployCmd.Flags().StringVar(&osName, "os", "", "")
	deployCmd.Flags().StringVar(&arch, "arch", "", "")
	deployCmd.Flags().StringVar(&binaryPath, "binary-path", "", "")

	deployCmd.MarkFlagRequired("token")
	deployCmd.MarkFlagRequired("org")
	deployCmd.MarkFlagRequired("name")
	deployCmd.MarkFlagRequired("pubkey-path")
	deployCmd.MarkFlagRequired("shasum-path")
	deployCmd.MarkFlagRequired("shasum-sig-path")
	deployCmd.MarkFlagRequired("version")
	deployCmd.MarkFlagRequired("os")
	deployCmd.MarkFlagRequired("arch")
	deployCmd.MarkFlagRequired("binary-path")
}

func DeployToTerraformCloud() (err error) {
	c := puver.NewTerraformCloudClient(token, org, providerName)

	file, err := os.Open(pubKeyPath)

	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = puver.CreateProvider(context.Background(), &c)

	if !errors.Is(err, puver.AlreadyExist) && err != nil {
		return err
	}

	keyId, err := puver.FindGPGKey(context.Background(), &c, string(content))

	if err != nil {
		return err
	}

	if keyId == "" {
		addGpgResponse, err := puver.AddGPGKey(context.Background(), &c, string(content))

		if err != nil {
			return err
		}

		keyId = addGpgResponse.KeyId
	}

	createVersionResponse, err := puver.CreateVersion(context.Background(), &c, version, keyId)

	if err != nil {
		if !errors.Is(err, puver.AlreadyExist) {
			return err
		}
	} else {
		for _, pair := range [][]string{
			{createVersionResponse.Data.Links.ShasumsSigUpload, shasumSigPath},
			{createVersionResponse.Data.Links.ShasumsUpload, shasumPath},
		} {
			err := puver.Upload(&c, pair[0], pair[1])
			if err != nil {
				return err
			}
		}
	}

	f, err := os.Open(shasumPath)

	if err != nil {
		return err
	}

	shasum, err := io.ReadAll(f)

	if err != nil {
		return err
	}

	hash := puver.HashData(string(shasum), osName, arch)

	if hash == "" {
		return errors.New("shasum empty")
	}

	createPlatformResponse, err := puver.CreatePlatform(context.Background(), &c, osName, arch, hash, filepath.Base(binaryPath), version)

	if err != nil {
		return err
	}

	puver.Upload(&c, createPlatformResponse.Data.Links.ProviderBinaryUpload, binaryPath)

	return nil
}
