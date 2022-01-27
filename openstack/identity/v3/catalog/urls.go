package catalog

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
