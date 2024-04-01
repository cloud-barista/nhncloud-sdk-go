// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2024.04

package subnets

import gophercloud "github.com/cloud-barista/nhncloud-sdk-go"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("subnets")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}
