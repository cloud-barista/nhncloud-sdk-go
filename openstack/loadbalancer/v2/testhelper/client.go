package common

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
