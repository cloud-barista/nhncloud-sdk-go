//go:build acceptance || baremetal || allocations
// +build acceptance baremetal allocations

package noauth

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	v1 "github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/openstack/baremetal/v1"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/baremetal/v1/allocations"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestAllocationsCreateDestroy(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewBareMetalV1NoAuthClient()
	th.AssertNoErr(t, err)

	client.Microversion = "1.52"

	allocation, err := v1.CreateAllocation(t, client)
	th.AssertNoErr(t, err)
	defer v1.DeleteAllocation(t, client, allocation)

	found := false
	err = allocations.List(client, allocations.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		allocationList, err := allocations.ExtractAllocations(page)
		if err != nil {
			return false, err
		}

		for _, a := range allocationList {
			if a.UUID == allocation.UUID {
				found = true
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, found, true)
}
