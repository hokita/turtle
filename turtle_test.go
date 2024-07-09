package main

import "testing"

func TestLoadConfig(t *testing.T) {
	yaml := []byte(`
endpoints:
  sample:
    url: "https://example.com"
`)

	config, _ := loadConfig(yaml)
	if config.Endpoints["sample"].URL != "https://example.com" {
		t.Fatalf("expected https://example.com, got %s", config.Endpoints["sample"].URL)
	}
}
