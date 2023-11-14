package messages

import "github.com/cloud-barista/nhncloud-sdk-go"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("messages")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("messages", id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return getURL(c, id)
}
