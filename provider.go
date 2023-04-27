package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func createProvider(ctx context.Context, c *TerraformCloudClient) error {
	b := CreateProviderRequest{
		CreateProviderData{
			Type: RegistryProviders,
			CreateProviderAttributes: CreateProviderAttributes{
				Name:         c.Name,
				Namespace:    c.Organization,
				RegistryName: RegistryVisibilityPrivate,
			},
		},
	}

	payload, err := json.Marshal(b)
	fmt.Println(string(payload))
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

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("[%d]: %s", resp.StatusCode, string(b))
	}

	return nil
}
