package puver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CreateProviderAttributes struct {
	Name         string             `json:"name"`
	Namespace    string             `json:"namespace"`
	RegistryName RegistryVisibility `json:"registry-name"`
}

type CreateProviderData struct {
	Type                     RequestType `json:"type"`
	CreateProviderAttributes `json:"attributes"`
}

type CreateProviderRequest struct {
	CreateProviderData `json:"data"`
}

func CreateProvider(ctx context.Context, c *TerraformCloudClient) error {
	b := CreateProviderRequest{
		CreateProviderData{
			Type: RequestTypeRegistryProviders,
			CreateProviderAttributes: CreateProviderAttributes{
				Name:         c.Name,
				Namespace:    c.Organization,
				RegistryName: RegistryVisibilityPrivate,
			},
		},
	}

	payload, err := json.Marshal(b)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://app.terraform.io/api/v2/organizations/%s/registry-providers", c.Organization)

	resp, err := c.Request(
		"POST",
		url,
		bytes.NewBuffer(payload),
		nil,
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusUnprocessableEntity && strings.Contains(string(body), "Name has already been taken") {
		return AlreadyExist
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("[%d]: %s", resp.StatusCode, string(body))
	}

	return nil
}
