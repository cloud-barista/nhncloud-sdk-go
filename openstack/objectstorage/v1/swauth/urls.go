package swauth

import "github.com/cloud-barista/nhncloud-sdk-go"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
