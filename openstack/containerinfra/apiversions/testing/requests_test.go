package testing

import (
	"fmt"
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-go/openstack/containerinfra/apiversions"
	th "github.com/cloud-barista/nhncloud-sdk-go/testhelper"
	"github.com/cloud-barista/nhncloud-sdk-go/testhelper/client"
)

func TestListAPIVersions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allVersions, err := apiversions.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := apiversions.ExtractAPIVersions(allVersions)
	th.AssertNoErr(t, err)
	fmt.Println(actual)
	th.AssertDeepEquals(t, MagnumAllAPIVersionResults, actual)
}

func TestGetAPIVersion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := apiversions.Get(client.ServiceClient(), "v1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, MagnumAPIVersion1Result, *actual)
}
