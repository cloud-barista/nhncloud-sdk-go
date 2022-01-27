package quotasets

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

const resourcePath = "os-quota-sets"

func getURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func getDefaultsURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID, "defaults")
}

func updateURL(c *gophercloud.ServiceClient, projectID string) string {
	return getURL(c, projectID)
}

func deleteURL(c *gophercloud.ServiceClient, projectID string) string {
	return getURL(c, projectID)
}
