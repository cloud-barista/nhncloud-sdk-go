//go:build acceptance || compute || limits
// +build acceptance compute limits

package v2

import (
	"testing"
	"time"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/compute/v2/extensions/instanceactions"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/compute/v2/servers"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestInstanceActions(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	allPages, err := instanceactions.List(client, server.ID, nil).AllPages()
	th.AssertNoErr(t, err)
	allActions, err := instanceactions.ExtractInstanceActions(allPages)
	th.AssertNoErr(t, err)

	var found bool

	for _, action := range allActions {
		action, err := instanceactions.Get(client, server.ID, action.RequestID).Extract()
		th.AssertNoErr(t, err)
		tools.PrintResource(t, action)

		if action.Action == "create" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestInstanceActionsMicroversions(t *testing.T) {
	clients.RequireLong(t)
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")
	clients.SkipRelease(t, "stable/queens")
	clients.SkipRelease(t, "stable/rocky")

	now := time.Now()

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)
	client.Microversion = "2.66"

	server, err := CreateMicroversionServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	rebootOpts := servers.RebootOpts{
		Type: servers.HardReboot,
	}

	err = servers.Reboot(client, server.ID, rebootOpts).ExtractErr()
	if err = WaitForComputeStatus(client, server, "ACTIVE"); err != nil {
		t.Fatal(err)
	}

	listOpts := instanceactions.ListOpts{
		Limit:        1,
		ChangesSince: &now,
	}

	allPages, err := instanceactions.List(client, server.ID, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allActions, err := instanceactions.ExtractInstanceActions(allPages)
	th.AssertNoErr(t, err)

	var found bool

	for _, action := range allActions {
		action, err := instanceactions.Get(client, server.ID, action.RequestID).Extract()
		th.AssertNoErr(t, err)
		tools.PrintResource(t, action)

		if action.Action == "reboot" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	listOpts = instanceactions.ListOpts{
		Limit:         1,
		ChangesBefore: &now,
	}

	allPages, err = instanceactions.List(client, server.ID, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allActions, err = instanceactions.ExtractInstanceActions(allPages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(allActions), 0)
}
