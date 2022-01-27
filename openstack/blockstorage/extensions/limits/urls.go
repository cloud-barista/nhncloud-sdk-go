package limits

import (
	"github.com/cloud-barista/nhncloud-sdk-for-drv"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
