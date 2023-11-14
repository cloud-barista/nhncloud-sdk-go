package availabilityzones

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}

// ### Note) NHN Cloud doesn't support 'ListDetail()' for availabilityzones!!
