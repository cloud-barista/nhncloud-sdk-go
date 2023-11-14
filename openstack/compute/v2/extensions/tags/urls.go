package tags

import "github.com/cloud-barista/nhncloud-sdk-go"

const (
	rootResourcePath = "servers"
	resourcePath     = "tags"
)

func rootURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL(rootResourcePath, serverID, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	return c.ServiceURL(rootResourcePath, serverID, resourcePath, tag)
}

func listURL(c *gophercloud.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}

func checkURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func replaceAllURL(c *gophercloud.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}

func addURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func deleteURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func deleteAllURL(c *gophercloud.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}
