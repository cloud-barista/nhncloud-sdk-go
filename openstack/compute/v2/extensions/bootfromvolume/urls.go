package bootfromvolume

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
