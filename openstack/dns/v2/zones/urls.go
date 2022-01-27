package zones

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func baseURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *gophercloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}
