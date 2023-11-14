package buildinfo

import "github.com/cloud-barista/nhncloud-sdk-go"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
