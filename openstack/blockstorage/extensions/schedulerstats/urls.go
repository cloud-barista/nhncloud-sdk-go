package schedulerstats

import "github.com/cloud-barista/nhncloud-sdk-go"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
