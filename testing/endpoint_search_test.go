package testing

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = gophercloud.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
