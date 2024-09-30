// Package clients contains functions for creating service clients
// for utils services.
// That clients can be used in acceptance tests.
package clients

import (
	"net/http"
	"os"

	"github.com/abramad-labs/gophercloud-utils-x/client"
	"github.com/abramad-labs/gophercloud-utils-x/env"
	"github.com/abramad-labs/gophercloud-utils-x/gnocchi"
	"github.com/abramad-labs/gophercloud-utils-x/openstack/clientconfig"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// NewGnocchiV1Client returns a *ServiceClient for making calls
// to the Gnocchi v1 API.
// An error will be returned if authentication or client
// creation was not possible.
func NewGnocchiV1Client() (*gophercloud.ServiceClient, error) {
	ao, err := clientconfig.AuthOptions(nil)
	if err != nil {
		return nil, err
	}

	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	if os.Getenv("OS_DEBUG") != "" {
		provider.HTTPClient = http.Client{
			Transport: &client.RoundTripper{
				Rt:     &http.Transport{},
				Logger: &client.DefaultLogger{},
			},
		}
	}

	err = openstack.Authenticate(provider, *ao)
	if err != nil {
		return nil, err
	}

	return gnocchi.NewGnocchiV1(provider, gophercloud.EndpointOpts{
		Region: env.Getenv("OS_REGION_NAME"),
	})
}
