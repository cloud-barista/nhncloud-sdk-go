package limits

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
