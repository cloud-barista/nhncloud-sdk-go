package groups

import "github.com/cloud-barista/nhncloud-sdk-go"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func getURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func updateURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func deleteURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}
