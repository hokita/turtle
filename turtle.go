package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

var ErrMissingTarget = errors.New("missing target argument")

func turtle() error {
	if len(os.Args) < 2 {
		return ErrMissingTarget
	}
	target := os.Args[1]

	resp, err := call(target)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

func call(target string) (string, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return "", err
	}

	config, err := loadConfig(yamlFile)
	if err != nil {
		return "", err
	}

	endpoint, exists := config.Endpoints[target]
	if !exists {
		return "", errors.New("endpoint not found")
	}

	resp, err := http.Get(endpoint.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func loadConfig(yamlFile []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	URL string `yaml:"url"`
}
