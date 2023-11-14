// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2022.07

package loadbalancers

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
	// "github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/listeners"
	// "github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/pools"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToLoadBalancerListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Loadbalancer attributes you want to see returned. SortKey allows you to
// sort by a particular attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.

type ListOpts struct {   									// Modified 
	ID                 string   `q:"id"`
	Name               string   `q:"name"`
	ProvisioningStatus string   `q:"provisioning_status"`
	Description        string   `q:"description"`
	VipAddress         string   `q:"vip_address"`
	VipPortID          string   `q:"vip_port_id"`
	VipSubnetID        string   `q:"vip_subnet_id"`
	OperatingStatus    string   `q:"operating_status"`
	LoadbalancerType   string   `q:"loadbalancer_type"`    // 'shared' or 'dedicated'
}

// ToLoadBalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// load balancers. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those load balancers that are owned by
// the project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToLoadBalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToLoadBalancerCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {   											// Modified 
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name 				string `json:"name,omitempty"`

	// Human-readable description for the Loadbalancer.
	Description 	 	string `json:"description,omitempty"`

	// The subnet on which to allocate the Loadbalancer's address. A project can
	// only create Loadbalancers on networks authorized by policy (e.g. networks
	// that belong to them or networks that are shared).
	VipSubnetID 	 	string `json:"vip_subnet_id" required:"true"`

	// The IP address of the Loadbalancer.
	VipAddress 		 	string `json:"vip_address,omitempty"`

	// The administrative state of the Loadbalancer. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp 	 	bool `json:"admin_state_up,omitempty"`
	
	LoadBalancerType 	string `json:"loadbalancer_type,omitempty"`	// 'shared' or 'dedicated'. (Added )
}

// ToLoadBalancerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "loadbalancer")
}

// Create is an operation which provisions a new loadbalancer based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToLoadBalancerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToLoadBalancerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {										// Modified 
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Human-readable description for the Loadbalancer.
	Description *string `json:"description,omitempty"`

	// The administrative state of the Loadbalancer. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToLoadBalancerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "loadbalancer")
}

// Update is an operation which modifies the attributes of the specified
// LoadBalancer.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToLoadBalancerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to the
// Delete request.
type DeleteOptsBuilder interface {
	ToLoadBalancerDeleteQuery() (string, error)
}

// DeleteOpts is the common options struct used in this package's Delete
// operation.
type DeleteOpts struct {
	// Cascade will delete all children of the load balancer (listners, monitors, etc).
	Cascade bool `q:"cascade"`
}

// ToLoadBalancerDeleteQuery formats a DeleteOpts into a query string.
func (opts DeleteOpts) ToLoadBalancerDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(c *gophercloud.ServiceClient, id string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := resourceURL(c, id)
	if opts != nil {
		query, err := opts.ToLoadBalancerDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	resp, err := c.Delete(url, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(c *gophercloud.ServiceClient, id string) (r GetStatusesResult) {
	resp, err := c.Get(statusRootURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetStats will return the shows the current statistics of a particular LoadBalancer.
func GetStats(c *gophercloud.ServiceClient, id string) (r StatsResult) {
	resp, err := c.Get(statisticsRootURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Failover performs a failover of a load balancer.
func Failover(c *gophercloud.ServiceClient, id string) (r FailoverResult) {
	resp, err := c.Put(failoverRootURL(c, id), nil, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
