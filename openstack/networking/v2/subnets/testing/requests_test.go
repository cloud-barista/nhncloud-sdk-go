package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/cloud-barista/nhncloud-sdk-go/openstack/networking/v2/common"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/networking/v2/subnets"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
	th "github.com/cloud-barista/nhncloud-sdk-go/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, SubnetListResult)
	})

	count := 0

	subnets.List(fake.ServiceClient(), subnets.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := subnets.ExtractSubnets(page)
		if err != nil {
			t.Errorf("Failed to extract subnets: %v", err)
			return false, nil
		}

		expected := []subnets.Subnet{
			Subnet1,
			Subnet2,
			Subnet3,
			Subnet4,
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
