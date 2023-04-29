package puver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreatePlatformAttributes struct {
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Shasum   string `json:"shasum"`
	Filename string `json:"filename"`
}

type CreatePlatformData struct {
	Type                     RequestType `json:"type"`
	CreatePlatformAttributes `json"attributes"`
}

type CreatePlatformRequest struct {
	CreatePlatformData `json:"data"`
}

type CreatePlatformResponse struct {
	Data struct {
		Links struct {
			ProviderBinaryUpload string `json:"provider-binary-upload"`
		} `json:"links"`
	} `json:"data"`
}

func CreatePlatform(
	ctx context.Context,
	c *TerraformCloudClient,
	os string,
	arch string,
	shasum string,
	filename string,
	version string,
) (CreatePlatformResponse, error) {
	var r CreatePlatformResponse

	b := CreatePlatformRequest{
		CreatePlatformData{
			Type: RequestTypeRegistryProviderVersionPlatforms,
			CreatePlatformAttributes: CreatePlatformAttributes{
				Os:       os,
				Arch:     arch,
				Shasum:   shasum,
				Filename: filename,
			},
		},
	}

	payload, err := json.Marshal(b)
	if err != nil {
		return r, err
	}

	url := fmt.Sprintf(
		"https://app.terraform.io/api/v2/organizations/%s/registry-providers/private/%s/%s/versions/%s/platforms",
		c.Organization,
		c.Organization,
		c.Name,
		version,
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
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return r, fmt.Errorf("[%d]: %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &r)

	if err != nil {
		return r, err
	}

	return r, nil
}
