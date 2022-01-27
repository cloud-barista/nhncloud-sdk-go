package stacktemplates

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func getURL(c *gophercloud.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "template")
}

func validateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("validate")
}
