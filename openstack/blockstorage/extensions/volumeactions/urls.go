package volumeactions

import "github.com/cloud-barista/nhncloud-sdk-go"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
