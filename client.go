package main

import (
	"fmt"
	"io"
	"net/http"
)

//go:generate enumer -type=RegistryVisibility -json -trimprefix RegistryVisibility -transform lower
type RegistryVisibility int

const (
	RegistryVisibilityPublic RegistryVisibility = iota
	RegistryVisibilityPrivate
)

//go:generate enumer -type=RequestType -json -transform kebab -trimprefix RequestType
type RequestType int

const (
	RequestTypeRegistryProviders RequestType = iota
	RequestTypeGPGKeys
)

func NewTerraformCloudClient(token string, org string, name string) TerraformCloudClient {
	return TerraformCloudClient{
		Token:        token,
		Organization: org,
		Name:         name,
	}
}

type TerraformCloudClient struct {
	Token        string
	Organization string
	Name         string
}

func (c TerraformCloudClient) defaultHeader() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", c.Token)
	headers["Content-Type"] = "application/vnd.api+json"
	return headers
}

func (c TerraformCloudClient) mergeHeader(h map[string]string) map[string]string {
	base := c.defaultHeader()
	for k, v := range h {
		base[k] = v
	}

	return base
}

func (c TerraformCloudClient) Request(method, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	for k, v := range c.mergeHeader(header) {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
