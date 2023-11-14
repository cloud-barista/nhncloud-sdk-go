//go:build acceptance
// +build acceptance

package v1

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-go/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/orchestration/v1/buildinfo"
	th "github.com/cloud-barista/nhncloud-sdk-go/testhelper"
)

func TestBuildInfo(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")

	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	bi, err := buildinfo.Get(client).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved build info: %+v\n", bi)
}
