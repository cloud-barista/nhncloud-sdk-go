package apiversions

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
)

func getURL(c *gophercloud.ServiceClient, version string) string {
	return c.ServiceURL(version)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL()
}
