// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2022.07

package subnets

import (
	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a subnet resource.
func (r commonResult) Extract() (*Subnet, error) {
	var s struct {
		Subnet *Subnet `json:"subnet"`
	}
	err := r.ExtractInto(&s)
	return s.Subnet, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Subnet.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Subnet.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Subnet.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// AllocationPool represents a sub-range of cidr available for dynamic
// allocation to ports, e.g. {Start: "10.0.0.2", End: "10.0.0.254"}
type AllocationPool struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// HostRoute represents a route that should be used by devices with IPs from
// a subnet (not including local subnet route).
type HostRoute struct {
	DestinationCIDR string `json:"destination"`
	NextHop         string `json:"nexthop"`
}

// Subnet represents a subnet. See package documentation for a top-level
// description of what this is.
type Subnet struct {											// Modified by B.T. Oh
	// Human-readable name for the subnet. Might not be unique.
	Name string `json:"name"`

	// Specifies whether DHCP is enabled for this subnet or not.
	EnableDHCP bool `json:"enable_dhcp"`

	// UUID of the parent network.
	NetworkID string `json:"network_id"`

	// TenantID is the project owner of the subnet.
	TenantID string `json:"tenant_id"`

	// DNS name servers used by hosts in this subnet.
	DNSNameservers []string `json:"dns_nameservers"`

	// Default gateway used by devices in this subnet.
	GatewayIP string `json:"gateway_ip"`

	// The IPv6 router advertisement specifies whether the networking service
	// should transmit ICMPv6 packets.
	IPv6RAMode string `json:"ipv6_ra_mode"`

	// Sub-ranges of CIDR available for dynamic allocation to ports.
	// See AllocationPool.
	AllocationPools []AllocationPool `json:"allocation_pools"`

	// Routes that should be used by devices with IPs from this subnet
	// (not including local subnet route).
	HostRoutes []HostRoute `json:"host_routes"`

	// IP version, either `4' or `6'.
	IPVersion int `json:"ip_version"`

	// The IPv6 address modes specifies mechanisms for assigning IPv6 IP addresses.
	IPv6AddressMode string `json:"ipv6_address_mode"`

	// CIDR representing IP range for this subnet, based on IP version.
	CIDR string `json:"cidr"`

	// UUID representing the subnet.
	ID string `json:"id"`

	// SubnetPoolID is the id of the subnet pool associated with the subnet.
	SubnetPoolID string `json:"subnetpool_id"`
}

// SubnetPage is the page returned by a pager when traversing over a collection
// of subnets.
type SubnetPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of subnets has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r SubnetPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"subnets_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SubnetPage struct is empty.
func (r SubnetPage) IsEmpty() (bool, error) {
	is, err := ExtractSubnets(r)
	return len(is) == 0, err
}

// ExtractSubnets accepts a Page struct, specifically a SubnetPage struct,
// and extracts the elements into a slice of Subnet structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractSubnets(r pagination.Page) ([]Subnet, error) {
	var s struct {
		Subnets []Subnet `json:"subnets"`
	}
	err := (r.(SubnetPage)).ExtractInto(&s)
	return s.Subnets, err
}
