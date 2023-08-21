package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Manifest struct {
	Version int
	Rules   []Values
}

type Values struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	TTL         string `yaml:"ttl"`
	MaxAttempts int    `yaml:"max_attempts"`
}

func LoadFile(path string) (*Manifest, error) {
	cleanedPath := filepath.Clean(path)
	file, err := os.ReadFile(cleanedPath)
	if err != nil {
		return nil, err
	}

	manifest := &Manifest{}

	if err = yaml.Unmarshal(file, manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}
