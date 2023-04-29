package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wim-web/puver/internal/puver"
)

var (
	name       string
	email      string
	passphrase string
)

var keyGenCmd = &cobra.Command{
	Use:   "key-gen",
	Short: "generate GPG Key",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := puver.GenerateGPGKey(
			name,
			email,
			passphrase,
		)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(key.Public)
		fmt.Printf("\n\n")
		fmt.Println(key.Private)
	},
}

func init() {
	rootCmd.AddCommand(keyGenCmd)

	keyGenCmd.Flags().StringVarP(&name, "name", "n", "", "")
	keyGenCmd.Flags().StringVarP(&email, "email", "e", "", "")
	keyGenCmd.Flags().StringVarP(&passphrase, "pass", "p", "", "")

	keyGenCmd.MarkFlagRequired("pass")
}
