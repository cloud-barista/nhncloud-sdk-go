package acls

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func containerURL(client *gophercloud.ServiceClient, containerID string) string {
	return client.ServiceURL("containers", containerID, "acl")
}

func secretURL(client *gophercloud.ServiceClient, secretID string) string {
	return client.ServiceURL("secrets", secretID, "acl")
}
