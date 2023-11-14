package endpointgroups

import "github.com/cloud-barista/nhncloud-sdk-go"

const (
	rootPath     = "vpn"
	resourcePath = "endpoint-groups"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
