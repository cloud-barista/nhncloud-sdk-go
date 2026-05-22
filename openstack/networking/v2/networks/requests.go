package networks

import (
	gophercloud "github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNetworkListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the network attributes you want to see returned. SortKey allows you to sort
// by a particular network attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
// https://docs.nhncloud.com/ko/Network/VPC/ko/public-api-compat/
type ListOpts struct {
	ID                  string `q:"id"`
	Name                string `q:"name"`
	ProviderNetworkType string `q:"provider:network_type"`
	RouterExternal      bool   `q:"router:external"`
	Shared              bool   `q:"shared"`
	TenantID            string `q:"tenant_id"`
	SortDir             string `q:"sort_dir"`
	SortKey             string `q:"sort_key"`
	Fields              string `q:"fields"` // Field Name of the network
}

// ToNetworkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNetworkListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// networks. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToNetworkListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return NetworkPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific network based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(getURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToNetworkCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a network.
type CreateOpts struct {
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Shared         *bool  `json:"shared,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	ProjectID      string `json:"project_id,omitempty"`
	RouterExternal *bool  `json:"router:external,omitempty"`
}

// ToNetworkCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToNetworkCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "network")
}

// Create accepts a CreateOpts struct and creates a new network using the values
// provided.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNetworkCreateMap()
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
	ToNetworkUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a network.
type UpdateOpts struct {
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Shared       *bool   `json:"shared,omitempty"`
}

// ToNetworkUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToNetworkUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "network")
}

// Update accepts an UpdateOpts struct and updates an existing network using the
// values provided.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNetworkUpdateMap()
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

// Delete accepts a unique ID and deletes the network associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
