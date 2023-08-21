package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileSuccess(t *testing.T) {
	// Create a temporary file.
	tmpFile, err := ioutil.TempFile("", "manifest*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write a valid YAML to the file.
	content := `
version: 1
rules:
  - name: rule1
    type: typeA
    ttl: 10s
  - name: rule2
    type: typeB
    ttl: 20s
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}

	// Load the file.
	manifest, err := LoadFile(tmpFile.Name())

	assert.NoError(t, err)
	assert.Equal(t, 1, manifest.Version)
	assert.Len(t, manifest.Rules, 2)
	assert.Equal(t, "rule1", manifest.Rules[0].Name)
	assert.Equal(t, "typeA", manifest.Rules[0].Type)
	assert.Equal(t, "10s", manifest.Rules[0].TTL)
	assert.Equal(t, "rule2", manifest.Rules[1].Name)
	assert.Equal(t, "typeB", manifest.Rules[1].Type)
	assert.Equal(t, "20s", manifest.Rules[1].TTL)
}

func TestLoadFileInvalidPath(t *testing.T) {
	_, err := LoadFile("/invalid/path/to/file.yaml")
	assert.Error(t, err)
}

func TestLoadFileInvalidYAML(t *testing.T) {
	// Create a temporary file.
	tmpFile, err := os.CreateTemp("", "manifest*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write invalid YAML content.
	content := `
version: 1
rules:
- name: rule1
type: typeA  # This line is mis-indented, making the YAML invalid.
ttl: 10s
- :
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}

	// Attempt to load the invalid YAML.
	_, err = LoadFile(tmpFile.Name())
	assert.Error(t, err)
}
