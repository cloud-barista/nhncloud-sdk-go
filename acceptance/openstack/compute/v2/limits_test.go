//go:build acceptance || compute || limits
// +build acceptance compute limits

package v2

import (
	"strings"
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-go/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-go/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/compute/v2/extensions/limits"
	th "github.com/cloud-barista/nhncloud-sdk-go/testhelper"
)

func TestLimits(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	limits, err := limits.Get(client, nil).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, limits)

	th.AssertEquals(t, limits.Absolute.MaxPersonalitySize, 10240)
}

func TestLimitsForTenant(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	// I think this is the easiest way to get the tenant ID while being
	// agnostic to Identity v2 and v3.
	// Technically we're just returning the limits for ourselves, but it's
	// the fact that we're specifying a tenant ID that is important here.
	endpointParts := strings.Split(client.Endpoint, "/")
	tenantID := endpointParts[4]

	getOpts := limits.GetOpts{
		TenantID: tenantID,
	}

	limits, err := limits.Get(client, getOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, limits)

	th.AssertEquals(t, limits.Absolute.MaxPersonalitySize, 10240)
}
