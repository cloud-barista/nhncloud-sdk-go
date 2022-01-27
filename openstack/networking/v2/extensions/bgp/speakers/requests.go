package speakers

import (
	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
)

// List the bgp speakers
func List(c *gophercloud.ServiceClient) pagination.Pager {
	url := listURL(c)
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return BGPSpeakerPage{pagination.SinglePageBase(r)}
	})
}

// Get retrieve the specific bgp speaker by its uuid
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
