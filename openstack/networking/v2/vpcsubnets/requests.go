// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2022.07

package vpcsubnets

import (
	gophercloud "github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToVpcsubnetListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the vpcsubnet attributes you want to see returned. SortKey allows you to sort
// by a particular vpcsubnet attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct { // Modified
	ID         string `q:"id"`
	Name       string `q:"name"`
	EnableDHCP *bool  `q:"enable_dhcp"`
	VPCID      string `q:"vpc_id"`
	CIDR       string `q:"cidr"`
	Shared     bool   `q:"shared"` // Whether the vpcsubnet is shared. (Added )
	SortDir    string `q:"sort_dir"`
	SortKey    string `q:"sort_key"`
	Fields     string `q:"fields"` // Field Name of the vpcsubnet
}

// ToVpcsubnetListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVpcsubnetListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// vpcsubnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those vpcsubnets that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToVpcsubnetListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return VpcsubnetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific vpcsubnet based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// List request.
type CreateOptsBuilder interface {
	ToVpcsubnetCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents the attributes used when creating a new vpcsubnet.
type CreateOpts struct {
	// NetworkID is the UUID of the network the vpcsubnet will be associated with.
	VpcID string `json:"vpc_id" required:"true"`

	// CIDR is the address CIDR of the vpcsubnet.
	CIDR string `json:"cidr,omitempty"`

	// Name is a human-readable name of the vpcsubnet.
	Name string `json:"name,omitempty"`

	// The UUID of the project who owns the Vpcsubnet. Only administrative users
	// can specify a project UUID other than their own.
	TenantID string `json:"tenant_id,omitempty"`
}

// ToVpcsubnetCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToVpcsubnetCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpcsubnet")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create accepts a CreateOpts struct and creates a new vpcsubnet using the values
// provided. You must remember to provide a valid NetworkID, CIDR and IP
// version.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVpcsubnetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(createURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVpcsubnetUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents the attributes used when updating an existing vpcsubnet.
type UpdateOpts struct {
	// Name is a human-readable name of the vpcsubnet.
	Name *string `json:"name,omitempty"`

	// Description of the vpcsubnet.
	Description *string `json:"description,omitempty"`

	// GatewayIP sets gateway information for the vpcsubnet. Setting to nil will
	// cause a default gateway to automatically be created. Setting to an empty
	// string will cause the vpcsubnet to be created with no gateway. Setting to
	// an explicit address will set that address as the gateway.
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// DNSNameservers are the nameservers to be set via DHCP.
	DNSNameservers *[]string `json:"dns_nameservers,omitempty"`

	// EnableDHCP will either enable to disable the DHCP service.
	EnableDHCP *bool `json:"enable_dhcp,omitempty"`
}

// ToVpcsubnetUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToVpcsubnetUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "vpcsubnet")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing vpcsubnet using the
// values provided.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVpcsubnetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(updateURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete accepts a unique ID and deletes the vpcsubnet associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
