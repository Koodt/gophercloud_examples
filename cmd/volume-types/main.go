package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"

	"github.com/koodt/gophercloud_examples/internal/auth"
)

func main() {
	// Load configuration from file
	config, err := auth.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Authenticate with OpenStack
	provider, err := auth.Authenticate(config)
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
