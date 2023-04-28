package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func Upload(
	c *TerraformCloudClient,
	uploadUrl string,
	filePath string,
) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	resp, err := c.Request(
		"PUT",
		uploadUrl,
		bytes.NewReader(fileContent),
		nil,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("[%d] %s", resp.StatusCode, string(body))
	}

	return nil
}
