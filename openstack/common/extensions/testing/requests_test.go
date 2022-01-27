package testing

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/common/extensions"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListExtensionsSuccessfully(t)

	count := 0

	extensions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := extensions.ExtractExtensions(page)
		th.AssertNoErr(t, err)
		th.AssertDeepEquals(t, ExpectedExtensions, actual)

		return true, nil
	})

	th.CheckEquals(t, 1, count)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetExtensionSuccessfully(t)

	actual, err := extensions.Get(client.ServiceClient(), "agent").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, SingleExtension, actual)
}
