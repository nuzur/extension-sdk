package config

import (
	"fmt"
	"path"

	"go.uber.org/config"
)

type ExtensionConfig struct {
	Metadata ExtensionConfigMetadata `yaml:"metadata"`
}

type ExtensionConfigMetadata struct {
	UUID        string `yaml:"uuid"`
	VersionUUID string `yaml:"version_uuid"`
	Identifier  string `yaml:"identifier"`
	DisplayName string `yaml:"display_name"`
	Description string `yaml:"description"`
	Steps       int32  `yaml:"steps"`
}

func New(pt string) (config.Provider, error) {

	yaml, err := config.NewYAML(
		config.File(path.Join(pt, "extension.yaml")),
	)

	if err != nil {
		fmt.Printf("error loading config: %v", err)
		return nil, err
	}

	return yaml, nil
}
