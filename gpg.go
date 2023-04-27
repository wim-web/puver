package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddGPGKeyAttributes struct {
	Namespace string `json:"namespace"`
	PublicKey string `json:"ascii-armor"`
	KeyId     string `json:"key-id,omitempty"`
}

type AddGPGKeyData struct {
	Type                RequestType `json:"type"`
	AddGPGKeyAttributes `json:"attributes"`
}

type AddGPGKeyRequest struct {
	AddGPGKeyData `json:"data"`
}

type AddGPGKEYResponse struct {
	AddGPGKeyData `json:"data"`
}

func AddGPGKey(ctx context.Context, c *TerraformCloudClient, pubKey string) (AddGPGKEYResponse, error) {
	var r AddGPGKEYResponse

	b := AddGPGKeyRequest{
		AddGPGKeyData{
			Type: RequestTypeGPGKeys,
			AddGPGKeyAttributes: AddGPGKeyAttributes{
				Namespace: c.Organization,
				PublicKey: pubKey,
			},
		},
	}

	payload, err := json.Marshal(b)
	if err != nil {
		return r, err
	}

	// fmt.Println(string(payload))

	url := fmt.Sprintf("https://app.terraform.io/api/registry/private/v2/gpg-keys")

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

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		return r, fmt.Errorf("[%d]: %s", resp.StatusCode, string(b))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return r, err
	}

	json.Unmarshal(body, &r)

	return r, nil
}
