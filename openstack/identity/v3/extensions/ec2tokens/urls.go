package ec2tokens

import "github.com/cloud-barista/nhncloud-sdk-go"

func ec2tokensURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ec2tokens")
}

func s3tokensURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("s3tokens")
}
