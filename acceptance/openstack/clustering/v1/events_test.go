//go:build acceptance || clustering || events
// +build acceptance clustering events

package v1

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/clustering/v1/events"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestEventsList(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	opts := events.ListOpts{
		Limit: 200,
	}

	allPages, err := events.List(client, opts).AllPages()
	th.AssertNoErr(t, err)

	allEvents, err := events.ExtractEvents(allPages)
	th.AssertNoErr(t, err)

	for _, event := range allEvents {
		tools.PrintResource(t, event)
	}
}
