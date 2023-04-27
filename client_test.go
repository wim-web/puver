package main

import (
	"reflect"
	"testing"
)

func TestTerraformCloudClient_defaultHeader(t *testing.T) {
	c := NewTerraformCloudClient("test-token")
	expected := map[string]string{
		"Authorization": "Bearer test-token",
		"Content-Type":  "application/vnd.api+json",
	}
	headers := c.defaultHeader()
	if !reflect.DeepEqual(headers, expected) {
		t.Errorf("defaultHeader() returned %+v, expected %+v", headers, expected)
	}
}

func TestTerraformCloudClient_mergeHeader(t *testing.T) {
	c := NewTerraformCloudClient("test-token")
	base := map[string]string{
		"Authorization": "Bearer test-token",
		"Content-Type":  "application/vnd.api+json",
	}

	// Test when base and input have no overlapping keys
	t.Run("Non-overlapping keys", func(t *testing.T) {
		input := map[string]string{
			"X-Custom-Header": "value",
		}
		expected := map[string]string{
			"Authorization":   "Bearer test-token",
			"Content-Type":    "application/vnd.api+json",
			"X-Custom-Header": "value",
		}
		merged := c.mergeHeader(input)
		if !reflect.DeepEqual(merged, expected) {
			t.Errorf("mergeHeader() returned %+v, expected %+v", merged, expected)
		}
	})

	// Test when input overwrites some keys in base
	t.Run("Overwriting keys", func(t *testing.T) {
		input := map[string]string{
			"Content-Type":    "application/json",
			"X-Custom-Header": "new-value",
		}
		expected := map[string]string{
			"Authorization":   "Bearer test-token",
			"Content-Type":    "application/json",
			"X-Custom-Header": "new-value",
		}
		merged := c.mergeHeader(input)
		if !reflect.DeepEqual(merged, expected) {
			t.Errorf("mergeHeader() returned %+v, expected %+v", merged, expected)
		}
	})

	// Make sure base was not modified
	if !reflect.DeepEqual(c.defaultHeader(), base) {
		t.Errorf("mergeHeader() modified the default headers")
	}
}
