package networks

import gophercloud "github.com/cloud-barista/nhncloud-sdk-go"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("networks")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}
