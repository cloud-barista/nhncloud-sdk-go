// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2022.07

package pools

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	// "github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/monitors"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPoolListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Pool attributes you want to see returned. SortKey allows you to
// sort by a particular Pool attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {										// Modified 
	ID             string `q:"id"`
	Name           string `q:"name"`
	LBMethod       string `q:"lb_algorithm"`
	Protocol       string `q:"protocol"`
	AdminStateUp   bool   `q:"admin_state_up"`
	MonitorID 	   string `q:"healthmonitor_id"`			// Added 
}

// ToPoolListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPoolListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// pools. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those pools that are owned by the
// project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToPoolListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PoolPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type LBMethod string
type Protocol string

// Supported attributes for create/update operations.
const (
	LBMethodRoundRobin       LBMethod = "ROUND_ROBIN"
	LBMethodLeastConnections LBMethod = "LEAST_CONNECTIONS"
	LBMethodSourceIp         LBMethod = "SOURCE_IP"

	ProtocolTCP   Protocol = "TCP"
	ProtocolUDP   Protocol = "UDP"
	ProtocolPROXY Protocol = "PROXY"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
	// Protocol PROXYV2 requires octavia microversion 2.22
	ProtocolPROXYV2 Protocol = "PROXYV2"
	// Protocol SCTP requires octavia microversion 2.23
	ProtocolSCTP Protocol = "SCTP"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPoolCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {												// Modified 
	// The Listener on which the members of the pool will be associated with.
	// Note: one of LoadbalancerID or ListenerID must be provided.
	ListenerID string `json:"listener_id" required:"true"`

	// The algorithm used to distribute load between the members of the pool. The
	// current specification supports LBMethodRoundRobin, LBMethodLeastConnections
	// and LBMethodSourceIp as valid values for this attribute.
	LBMethod LBMethod `json:"lb_algorithm" required:"true"`

	// The protocol used by the pool members, you can use either
	// ProtocolTCP, ProtocolUDP, ProtocolPROXY, ProtocolHTTP, ProtocolHTTPS,
	// ProtocolSCTP or ProtocolPROXYV2.
	Protocol Protocol `json:"protocol" required:"true"`

	// Human-readable description for the pool.
	Description string `json:"description,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// Member's port for receiving. Deliver traffic to this port. The default value is -1.
	MemberPort int `json:"member_port,omitempty"`	  					// Added 

	// Persistence is the session persistence of the pool.
	// Omit this field to prevent session persistence.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`

	// Name of the pool.
	Name string `json:"name,omitempty"`
}

// ToPoolCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPoolCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "pool")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// load balancer pool.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPoolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular pool based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPoolUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {													// Modified 
	// The algorithm used to distribute load between the members of the pool. The
	// current specification supports LBMethodRoundRobin, LBMethodLeastConnections
	// and LBMethodSourceIp as valid values for this attribute.
	LBMethod LBMethod `json:"lb_algorithm,omitempty"`

	// Human-readable description for the pool.
	Description *string `json:"description,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Persistence is the session persistence of the pool.
	// Omit this field to prevent session persistence.
	Persistence *SessionPersistence `json:"session_persistence,omitempty"`  // Added 

	// Name of the pool.
	Name *string `json:"name,omitempty"`
}

// ToPoolUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPoolUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "pool")
}

// Update allows pools to be updated.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoolUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete will permanently delete a particular pool based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(resourceURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListMemberOptsBuilder allows extensions to add additional parameters to the
// ListMembers request.
type ListMembersOptsBuilder interface {
	ToMembersListQuery() (string, error)
}

// ListMembersOpts allows the filtering and sorting of paginated collections
// through the API. Filtering is achieved by passing in struct field values
// that map to the Member attributes you want to see returned. SortKey allows
// you to sort by a particular Member attribute. SortDir sets the direction,
// and is either `asc' or `desc'. Marker and Limit are used for pagination.
type ListMembersOpts struct {								// Modified 
	ID           	string `q:"id"`
	Weight       	int    `q:"weight"`
	AdminStateUp 	*bool  `q:"admin_state_up"`
	SubnetID     	string `q:"subnet_id"`					// Added 
	TenantID     	string `q:"tenant_id"`					// Added 
	Address      	string `q:"address"`
	ProtocolPort 	int    `q:"protocol_port"`
	OperatingStatus	string `q:"operating_status	"`			// Added 
}

// ToMemberListQuery formats a ListOpts into a query string.
func (opts ListMembersOpts) ToMembersListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListMembers returns a Pager which allows you to iterate over a collection of
// members. It accepts a ListMembersOptsBuilder, which allows you to filter and
// sort the returned collection for greater efficiency.
//
// Default policy settings return only those members that are owned by the
// project who submits the request, unless an admin user submits the request.
func ListMembers(c *gophercloud.ServiceClient, poolID string, opts ListMembersOptsBuilder) pagination.Pager {
	url := memberRootURL(c, poolID)
	if opts != nil {
		query, err := opts.ToMembersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return MemberPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateMemberOptsBuilder allows extensions to add additional parameters to the
// CreateMember request.
type CreateMemberOptsBuilder interface {
	ToMemberCreateMap() (map[string]interface{}, error)
}

// CreateMemberOpts is the common options struct used in this package's CreateMember
// operation.
type CreateMemberOpts struct {									// Modified 
	// A positive integer value that indicates the relative portion of traffic
	// that this member should receive from the pool. For example, a member with
	// a weight of 10 receives five times as much traffic as a member with a
	// weight of 2.
	Weight int `json:"weight,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp bool `json:"admin_state_up,omitempty"`

	// If you omit this parameter, LBaaS uses the vip_subnet_id parameter value
	// for the subnet UUID.
	SubnetID string `json:"subnet_id" required:"true"`

	// The IP address of the member to receive traffic from the load balancer.
	Address string `json:"address" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`
}

// ToMemberCreateMap builds a request body from CreateMemberOpts.
func (opts CreateMemberOpts) ToMemberCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "member")
}

// CreateMember will create and associate a Member with a particular Pool.
func CreateMember(c *gophercloud.ServiceClient, poolID string, opts CreateMemberOptsBuilder) (r CreateMemberResult) {
	b, err := opts.ToMemberCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(memberRootURL(c, poolID), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// GetMember retrieves a particular Pool Member based on its unique ID.
func GetMember(c *gophercloud.ServiceClient, poolID string, memberID string) (r GetMemberResult) {
	resp, err := c.Get(memberResourceURL(c, poolID, memberID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateMemberOptsBuilder allows extensions to add additional parameters to the
// List request.
type UpdateMemberOptsBuilder interface {
	ToMemberUpdateMap() (map[string]interface{}, error)
}

// UpdateMemberOpts is the common options struct used in this package's Update
// operation.
type UpdateMemberOpts struct {
	// A positive integer value that indicates the relative portion of traffic
	// that this member should receive from the pool. For example, a member with
	// a weight of 10 receives five times as much traffic as a member with a
	// weight of 2.
	Weight *int `json:"weight,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToMemberUpdateMap builds a request body from UpdateMemberOpts.
func (opts UpdateMemberOpts) ToMemberUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "member")
}

// Update allows Member to be updated.
func UpdateMember(c *gophercloud.ServiceClient, poolID string, memberID string, opts UpdateMemberOptsBuilder) (r UpdateMemberResult) {
	b, err := opts.ToMemberUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(memberResourceURL(c, poolID, memberID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// BatchUpdateMemberOptsBuilder allows extensions to add additional parameters to the BatchUpdateMembers request.
type BatchUpdateMemberOptsBuilder interface {
	ToBatchMemberUpdateMap() (map[string]interface{}, error)
}

// BatchUpdateMemberOpts is the common options struct used in this package's BatchUpdateMembers
// operation.
type BatchUpdateMemberOpts struct {
	// The IP address of the member to receive traffic from the load balancer.
	Address string `json:"address" required:"true"`

	// The port on which to listen for client traffic.
	ProtocolPort int `json:"protocol_port" required:"true"`

	// Name of the Member.
	Name *string `json:"name,omitempty"`

	// ProjectID is the UUID of the project who owns the Member.
	// Only administrative users can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// A positive integer value that indicates the relative portion of traffic
	// that this member should receive from the pool. For example, a member with
	// a weight of 10 receives five times as much traffic as a member with a
	// weight of 2.
	Weight *int `json:"weight,omitempty"`

	// If you omit this parameter, LBaaS uses the vip_subnet_id parameter value
	// for the subnet UUID.
	SubnetID *string `json:"subnet_id,omitempty"`

	// The administrative state of the Pool. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Is the member a backup? Backup members only receive traffic when all
	// non-backup members are down.
	// Requires microversion 2.1 or later.
	Backup *bool `json:"backup,omitempty"`

	// An alternate IP address used for health monitoring a backend member.
	MonitorAddress *string `json:"monitor_address,omitempty"`

	// An alternate protocol port used for health monitoring a backend member.
	MonitorPort *int `json:"monitor_port,omitempty"`

	// A list of simple strings assigned to the resource.
	// Requires microversion 2.5 or later.
	Tags []string `json:"tags,omitempty"`
}

// ToBatchMemberUpdateMap builds a request body from BatchUpdateMemberOpts.
func (opts BatchUpdateMemberOpts) ToBatchMemberUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if b["subnet_id"] == "" {
		b["subnet_id"] = nil
	}

	return b, nil
}

// BatchUpdateMembers updates the pool members in batch
func BatchUpdateMembers(c *gophercloud.ServiceClient, poolID string, opts []BatchUpdateMemberOpts) (r UpdateMembersResult) {
	members := []map[string]interface{}{}
	for _, opt := range opts {
		b, err := opt.ToBatchMemberUpdateMap()
		if err != nil {
			r.Err = err
			return
		}
		members = append(members, b)
	}

	b := map[string]interface{}{"members": members}

	resp, err := c.Put(memberRootURL(c, poolID), b, nil, &gophercloud.RequestOpts{OkCodes: []int{202}})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteMember will remove and disassociate a Member from a particular Pool.
func DeleteMember(c *gophercloud.ServiceClient, poolID string, memberID string) (r DeleteMemberResult) {
	resp, err := c.Delete(memberResourceURL(c, poolID, memberID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
