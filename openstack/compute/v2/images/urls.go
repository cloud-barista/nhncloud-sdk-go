package images

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

// Added by BT.OH
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images", "")
}

func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}
