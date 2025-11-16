package auth

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// Authenticate creates an authenticated OpenStack provider client
func Authenticate(config *Config) (*gophercloud.ProviderClient, error) {
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

