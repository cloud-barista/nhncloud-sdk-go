package tokens

import "github.com/cloud-barista/nhncloud-sdk-go"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
