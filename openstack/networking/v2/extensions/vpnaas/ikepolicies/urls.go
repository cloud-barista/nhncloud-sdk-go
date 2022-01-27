package ikepolicies

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

const (
	rootPath     = "vpn"
	resourcePath = "ikepolicies"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
