//go:build acceptance || networking || fwaas
// +build acceptance networking fwaas

package vpnaas

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	layer3 "github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/openstack/networking/v2/extensions/layer3"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/networking/v2/extensions/vpnaas/services"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestServiceList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	allPages, err := services.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}

func TestServiceCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	router, err := layer3.CreateExternalRouter(t, client)
	th.AssertNoErr(t, err)
	defer layer3.DeleteRouter(t, client, router.ID)

	service, err := CreateService(t, client, router.ID)
	th.AssertNoErr(t, err)
	defer DeleteService(t, client, service.ID)

	newService, err := services.Get(client, service.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, service)
	tools.PrintResource(t, newService)
}
