//go:build acceptance || networking || networkipavailabilities
// +build acceptance networking networkipavailabilities

package networkipavailabilities

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/networking/v2/extensions/networkipavailabilities"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestNetworkIPAvailabilityList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	allPages, err := networkipavailabilities.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allAvailabilities, err := networkipavailabilities.ExtractNetworkIPAvailabilities(allPages)
	th.AssertNoErr(t, err)

	for _, availability := range allAvailabilities {
		for _, subnet := range availability.SubnetIPAvailabilities {
			tools.PrintResource(t, subnet)
			tools.PrintResource(t, subnet.TotalIPs)
			tools.PrintResource(t, subnet.UsedIPs)
		}
	}
}
