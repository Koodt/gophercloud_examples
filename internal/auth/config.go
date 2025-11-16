package auth

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents OpenStack authentication configuration
type Config struct {
	AuthURL     string `yaml:"auth_url"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	ProjectName string `yaml:"project_name"`
	DomainName  string `yaml:"domain_name"`
	Region      string `yaml:"region"`
}

// LoadConfig loads OpenStack configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
