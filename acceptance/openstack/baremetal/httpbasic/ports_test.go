//go:build acceptance || baremetal || ports
// +build acceptance baremetal ports

package httpbasic

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	v1 "github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/openstack/baremetal/v1"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/baremetal/v1/ports"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"

	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestPortsCreateDestroy(t *testing.T) {
	clients.RequireLong(t)
	clients.RequireIronicHTTPBasic(t)

	client, err := clients.NewBareMetalV1HTTPBasic()
	th.AssertNoErr(t, err)
	client.Microversion = "1.53"

	node, err := v1.CreateFakeNode(t, client)
	port, err := v1.CreatePort(t, client, node)
	th.AssertNoErr(t, err)
	defer v1.DeleteNode(t, client, node)
	defer v1.DeletePort(t, client, port)

	found := false
	err = ports.List(client, ports.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		portList, err := ports.ExtractPorts(page)
		if err != nil {
			return false, err
		}

		for _, p := range portList {
			if p.UUID == port.UUID {
				found = true
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, found, true)
}

func TestPortsUpdate(t *testing.T) {
	clients.RequireLong(t)
	clients.RequireIronicHTTPBasic(t)

	client, err := clients.NewBareMetalV1HTTPBasic()
	th.AssertNoErr(t, err)
	client.Microversion = "1.53"

	node, err := v1.CreateFakeNode(t, client)
	port, err := v1.CreatePort(t, client, node)
	th.AssertNoErr(t, err)
	defer v1.DeleteNode(t, client, node)
	defer v1.DeletePort(t, client, port)

	updated, err := ports.Update(client, port.UUID, ports.UpdateOpts{
		ports.UpdateOperation{
			Op:    ports.ReplaceOp,
			Path:  "/address",
			Value: "aa:bb:cc:dd:ee:ff",
		},
	}).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, updated.Address, "aa:bb:cc:dd:ee:ff")
}
