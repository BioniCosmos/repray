package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

func FromArgs() ([]Config, error) {
	raw, err := parseArgs()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse command-line arguments: %w", err)
	}
	config, err := parseConfig(raw)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse the configuration: %w", err)
	}
	return config, nil
}

func parseArgs() ([]byte, error) {
	configPath := os.Args[1]
	return os.ReadFile(configPath)
}

type Config struct {
	Listen   string
	Upstream *url.URL
	TLS      *TLS
}

type TLS struct {
	CertFile string
	KeyFile  string
}

func parseConfig(raw []byte) ([]Config, error) {
	type configJson struct {
		Listen   string
		Upstream string
		TLS      *TLS
	}

	parsedJson := make([]configJson, 0)
	if err := json.Unmarshal(raw, &parsedJson); err != nil {
		return nil, err
	}

	parsedConfig := make([]Config, 0, len(parsedJson))
	for _, c := range parsedJson {
		upstream, err := url.Parse(c.Upstream)
		if err != nil {
			return nil, err
		}
		parsedConfig = append(parsedConfig, Config{Listen: c.Listen, Upstream: upstream, TLS: c.TLS})
	}
	return parsedConfig, nil
}
