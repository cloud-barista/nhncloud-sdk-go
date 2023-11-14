package testing

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-go/openstack/compute/apiversions"
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

	th.AssertDeepEquals(t, NovaAllAPIVersionResults, actual)
}

func TestGetAPIVersion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := apiversions.Get(client.ServiceClient(), "v2.1").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, NovaAPIVersion21Result, *actual)
}

func TestGetMultipleAPIVersion(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMultipleResponses(t)

	_, err := apiversions.Get(client.ServiceClient(), "v3").Extract()
	th.AssertEquals(t, err.Error(), "Unable to find requested API version")
}
