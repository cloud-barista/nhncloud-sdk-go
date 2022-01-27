package accounts

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func getURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *gophercloud.ServiceClient) string {
	return getURL(c)
}
