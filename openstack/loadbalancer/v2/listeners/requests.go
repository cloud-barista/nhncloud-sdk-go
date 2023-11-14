// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2022.07

package listeners

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
	// "github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/l7policies"
	// "github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/pools"
)

// Type Protocol represents a listener protocol.
type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP   Protocol = "TCP"
	ProtocolUDP   Protocol = "UDP"
	ProtocolPROXY Protocol = "PROXY"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
	// Protocol SCTP requires octavia microversion 2.23
	ProtocolSCTP            Protocol = "SCTP"
	ProtocolTerminatedHTTPS Protocol = "TERMINATED_HTTPS"
)

// Type TLSVersion represents a tls version
type TLSVersion string

const (
	TLSVersionSSLv3   TLSVersion = "SSLv3"
	TLSVersionTLSv1   TLSVersion = "TLSv1"
	TLSVersionTLSv1_1 TLSVersion = "TLSv1.1"
	TLSVersionTLSv1_2 TLSVersion = "TLSv1.2"
	TLSVersionTLSv1_3 TLSVersion = "TLSv1.3"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular listener attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {    									// Modified 
	DefaultPoolID        string `q:"default_pool_id"`
	Protocol             string `q:"protocol"`
	Description          string `q:"description"`
	Name                 string `q:"name"`
	AdminStateUp         bool   `q:"admin_state_up"`
	ConnectionLimit      int    `q:"connection_limit"`
	KeepAliveTimeout 	 int    `q:"keep_alive_timeout"`
	ProtocolPort         int    `q:"protocol_port"`
	ID                   string `q:"id"`
}

// ToListenerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// listeners. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those listeners that are owned by the
// project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options for creating a listener.
type CreateOpts struct {   									// Modified 
	// The protocol - can either be TCP, SCTP, HTTP, HTTPS or TERMINATED_HTTPS.
	Protocol 				Protocol `json:"protocol" required:"true"`

	// Human-readable description for the Listener.
	Description 			string `json:"description,omitempty"`
	
	// Human-readable name for the Listener. Does not have to be unique.
	Name 					string `json:"name,omitempty"`

	// The load balancer on which to provision this listener.
	LoadbalancerID 			string `json:"loadbalancer_id" required:"true"`

	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp 		   	bool `json:"admin_state_up,omitempty"`

	// The maximum number of connections allowed for the Listener.
	ConnLimit 			   	int `json:"connection_limit,omitempty"`

	KeepAliveTimeout 	   	int `json:"keepalive_timeout,omitempty"`

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef 	string `json:"default_tls_container_ref,omitempty"`

	// A list of references to TLS secrets.
	SniContainerRefs 	   	[]string `json:"sni_container_refs,omitempty"`

	// The port on which to listen for client traffic.
	ProtocolPort 		   	int `json:"protocol_port" required:"true"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "listener")
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other projects by
// specifying a ProjectID attribute different than their own.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular Listeners based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options for updating a Listener.
type UpdateOpts struct {													// Modified 
	// Human-readable description for the Listener.
	Description *string `json:"description,omitempty"`

	// Human-readable name for the Listener. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// The maximum number of connections allowed for the Listener.
	ConnLimit *int `json:"connection_limit,omitempty"`

	// Frontend client inactivity timeout in milliseconds
	KeepAliveTimeout *int `json:"keepalive_timeout,omitempty"`				// Added 

	// A reference to a Barbican container of TLS secrets.
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	// A list of references to TLS secrets.
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`
}

// ToListenerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "listener")
	if err != nil {
		return nil, err
	}

	if m := b["listener"].(map[string]interface{}); m["default_pool_id"] == "" {
		m["default_pool_id"] = nil
	}

	return b, nil
}

// Update is an operation which modifies the attributes of the specified
// Listener.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
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

// Delete will permanently delete a particular Listeners based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(resourceURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetStats will return the shows the current statistics of a particular Listeners.
func GetStats(c *gophercloud.ServiceClient, id string) (r StatsResult) {
	resp, err := c.Get(statisticsRootURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
