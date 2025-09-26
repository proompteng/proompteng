package main

import (
	"errors"
	"os"
	"strings"

	syaml "sigs.k8s.io/yaml"
)

type operatorConfig struct {
	LogLevel     string          `yaml:"logLevel"`
	FeatureFlags map[string]bool `yaml:"featureFlags"`
}

func loadOperatorConfig() (operatorConfig, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "/config/config.yaml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return operatorConfig{}, nil
		}
		return operatorConfig{}, err
	}

	var cfg operatorConfig
	if err := syaml.Unmarshal(data, &cfg); err != nil {
		return operatorConfig{}, err
	}

	return cfg, nil
}

func isDebugEnabled(envLevel string, cfg operatorConfig) bool {
	if strings.EqualFold(envLevel, "debug") {
		return true
	}
	return strings.EqualFold(cfg.LogLevel, "debug")
}
