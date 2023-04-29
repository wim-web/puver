package puver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CreateVersionAttributes struct {
	Version   string     `json:"version"`
	KeyId     string     `json:"key-id"`
	Protocols []Protocol `json:"protocols"`
}

type CreateVersionData struct {
	Type                    RequestType `json:"type"`
	CreateVersionAttributes `json:"attributes"`
}

type CreateVersionRequest struct {
	CreateVersionData `json:"data"`
}

type CreateVersionResponse struct {
	Data struct {
		Links struct {
			ShasumsUpload    string `json:"shasums-upload"`
			ShasumsSigUpload string `json:"shasums-sig-upload"`
		} `json:"links"`
	} `json:"data"`
}

func CreateVersion(
	ctx context.Context,
	c *TerraformCloudClient,
	version string,
	keyId string,
) (CreateVersionResponse, error) {
	var r CreateVersionResponse

	b := CreateVersionRequest{
		CreateVersionData{
			Type: RequestTypeRegistryProviderVersions,
			CreateVersionAttributes: CreateVersionAttributes{
				Version: version,
				KeyId:   keyId,
				Protocols: []Protocol{
					Protocol_6_0,
				},
			},
		},
	}

	payload, err := json.Marshal(b)
	if err != nil {
		return r, err
	}

	url := fmt.Sprintf(
		"https://app.terraform.io/api/v2/organizations/%s/registry-providers/private/%s/%s/versions",
		c.Organization,
		c.Organization,
		c.Name,
	)

	resp, err := c.Request(
		"POST",
		url,
		bytes.NewBuffer(payload),
		nil,
	)

	if err != nil {
		return r, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return r, err
	}

	if resp.StatusCode == http.StatusUnprocessableEntity && strings.Contains(string(body), "Version has already been taken") {
		return r, AlreadyExist
	}

	if resp.StatusCode != http.StatusCreated {
		return r, fmt.Errorf("[%d]: %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &r)

	if err != nil {
		return r, err
	}

	return r, nil
}
