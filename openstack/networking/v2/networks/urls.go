package networks

import gophercloud "github.com/cloud-barista/nhncloud-sdk-go"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("networks")
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("networks", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
