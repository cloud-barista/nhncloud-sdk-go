package apiversions

import (
	"strings"

	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}
