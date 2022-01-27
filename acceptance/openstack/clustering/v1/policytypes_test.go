//go:build acceptance || clustering || policytypes
// +build acceptance clustering policytypes

package v1

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/clustering/v1/policytypes"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestPolicyTypeList(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	allPages, err := policytypes.List(client).AllPages()
	th.AssertNoErr(t, err)

	allPolicyTypes, err := policytypes.ExtractPolicyTypes(allPages)
	th.AssertNoErr(t, err)

	for _, v := range allPolicyTypes {
		tools.PrintResource(t, v)
	}
}

func TestPolicyTypeList_v_1_5(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	client.Microversion = "1.5"
	allPages, err := policytypes.List(client).AllPages()
	th.AssertNoErr(t, err)

	allPolicyTypes, err := policytypes.ExtractPolicyTypes(allPages)
	th.AssertNoErr(t, err)

	for _, v := range allPolicyTypes {
		tools.PrintResource(t, v)
	}
}

func TestPolicyTypeGet(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	policyType, err := policytypes.Get(client, "senlin.policy.batch-1.0").Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, policyType)
}

func TestPolicyTypeGet_v_1_5(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	client.Microversion = "1.5"
	policyType, err := policytypes.Get(client, "senlin.policy.batch-1.0").Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, policyType)
}
