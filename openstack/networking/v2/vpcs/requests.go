// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Created by ETRI, 2024.02

package vpcs

import (
	"fmt"
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToVPCListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the network attributes you want to see returned. SortKey allows you to sort
// by a particular network attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Status       string `q:"status"`
	Name         string `q:"name"`
	Description  string `q:"description"`
	AdminStateUp *bool  `q:"admin_state_up"`
	TenantID     string `q:"tenant_id"`
	ProjectID    string `q:"project_id"`
	Shared       *bool  `q:"shared"`
	ID           string `q:"id"`
	Marker       string `q:"marker"`
	Limit        int    `q:"limit"`
	SortKey      string `q:"sort_key"`
	SortDir      string `q:"sort_dir"`
	Tags         string `q:"tags"`
	TagsAny      string `q:"tags-any"`
	NotTags      string `q:"not-tags"`
	NotTagsAny   string `q:"not-tags-any"`
}

// ToVPCListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVPCListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// networks. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToVPCListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	fmt.Printf("\n### Call URL : %s\n\n", url)

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return VPCPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific network based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	fmt.Printf("\n### Call URL : %s\n\n", getURL(c, id))

	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVPCCreateMap() (map[string]interface{}, error)
}

type NewVPC struct {
	Name   				string  	`json:"name"`
	CIDRv4				string  	`json:"cidrv4"`
	TenantID			string 		`json:"tenant_id,omitempty"`
	ExternalNetworkID 	string		`json:"external_network_id,omitempty"`
}

// CreateOpts represents options used to create a network.
type CreateOpts struct {
	VPC NewVPC `json:"vpc"`
}

// ToVPCCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToVPCCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "vpc") // Caution!!
}

// Create accepts a CreateOpts struct and creates a new network using the values
// provided. This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// network. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVPCCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	fmt.Printf("\n### Map : %v\n", b)
	fmt.Printf("\n### Body : %v\n\n", r)

	resp, err := c.Post(createURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVPCUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a network.
type UpdateOpts struct {
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Shared       *bool   `json:"shared,omitempty"`
}

// ToVPCUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToVPCUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "vpc")
}

// Update accepts a UpdateOpts struct and updates an existing network using the
// values provided. For more information, see the Create function.
func Update(c *gophercloud.ServiceClient, networkID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVPCUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(updateURL(c, networkID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete accepts a unique ID and deletes the network associated with it.
func Delete(c *gophercloud.ServiceClient, networkID string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, networkID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
