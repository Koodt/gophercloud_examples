package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
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

func main() {
	// Load configuration from file
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Authenticate with OpenStack
	provider, err := authenticate(config)
	if err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	// Get Block Storage service client
	client, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: config.Region,
		Type:   "block-storage",
	})
	if err != nil {
		log.Fatalf("Failed to create block storage client: %v", err)
	}

	// List volume types
	fmt.Println("Fetching volume types...")
	allPages, err := volumetypes.List(client, volumetypes.ListOpts{}).AllPages()
	if err != nil {
		log.Fatalf("Failed to list volume types: %v", err)
	}

	volumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		log.Fatalf("Failed to extract volume types: %v", err)
	}

	// Display volume types
	fmt.Printf("\nFound %d volume type(s):\n\n", len(volumeTypes))
	for i, vt := range volumeTypes {
		fmt.Printf("Volume Type #%d:\n", i+1)
		fmt.Printf("  ID: %s\n", vt.ID)
		fmt.Printf("  Name: %s\n", vt.Name)
		fmt.Printf("  Description: %s\n", vt.Description)
		fmt.Printf("  Is Public: %v\n", vt.IsPublic)
		fmt.Printf("  Extra Specs: %v\n", vt.ExtraSpecs)
		fmt.Println()
	}

	// Optionally, print as JSON
	if len(os.Args) > 1 && os.Args[1] == "--json" {
		jsonOutput, err := json.MarshalIndent(volumeTypes, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal JSON: %v", err)
		}
		fmt.Println("JSON output:")
		fmt.Println(string(jsonOutput))
	}
}

func loadConfig(filename string) (*Config, error) {
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

func authenticate(config *Config) (*gophercloud.ProviderClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: config.AuthURL,
		Username:         config.Username,
		Password:         config.Password,
		TenantName:       config.ProjectName,
		DomainName:       config.DomainName,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	return provider, nil
}
