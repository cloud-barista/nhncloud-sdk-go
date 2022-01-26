package tokens

import "github.com/cloud-barista/nhncloud-sdk-for-drv"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
