package serviceassets

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
