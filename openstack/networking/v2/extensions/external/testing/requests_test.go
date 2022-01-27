package testing

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/networking/v2/extensions/external"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/networking/v2/networks"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestListExternal(t *testing.T) {
	var iTrue bool = true

	networkListOpts := networks.ListOpts{
		ID: "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	}

	listOpts := external.ListOptsExt{
		ListOptsBuilder: networkListOpts,
		External:        &iTrue,
	}

	actual, err := listOpts.ToNetworkListQuery()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ExpectedListOpts, actual)
}
