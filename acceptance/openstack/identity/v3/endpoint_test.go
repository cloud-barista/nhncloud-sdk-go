//go:build acceptance
// +build acceptance

package v3

import (
	"strings"
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/identity/v3/endpoints"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/identity/v3/services"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestEndpointsList(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := endpoints.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, endpoint := range allEndpoints {
		tools.PrintResource(t, endpoint)

		if strings.Contains(endpoint.URL, "/v3") {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestEndpointsNavigateCatalog(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	// Discover the service we're interested in.
	serviceListOpts := services.ListOpts{
		ServiceType: "compute",
	}

	allPages, err := services.List(client, serviceListOpts).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(allServices), 1)

	computeService := allServices[0]
	tools.PrintResource(t, computeService)

	// Enumerate the endpoints available for this service.
	endpointListOpts := endpoints.ListOpts{
		Availability: gophercloud.AvailabilityPublic,
		ServiceID:    computeService.ID,
	}

	allPages, err = endpoints.List(client, endpointListOpts).AllPages()
	th.AssertNoErr(t, err)

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(allServices), 1)

	tools.PrintResource(t, allEndpoints[0])
}
