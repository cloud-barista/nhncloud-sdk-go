package portforwarding

import "github.com/cloud-barista/nhncloud-sdk-go"

const resourcePath = "floatingips"
const portForwardingPath = "port_forwardings"

func portForwardingUrl(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, portForwardingPath)
}

func singlePortForwardingUrl(c *gophercloud.ServiceClient, id string, portForwardingID string) string {
	return c.ServiceURL(resourcePath, id, portForwardingPath, portForwardingID)
}
