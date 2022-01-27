package buildinfo

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
