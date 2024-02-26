// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Modified by ETRI, 2024.02

package subnets

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a subnet resource.
func (r commonResult) Extract() (*Subnet, error) {
	var s struct {
		Subnet *Subnet `json:"vpcsubnet"` // Modified
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

type Subnet struct {
    RouterExternal    bool         `json:"router:external"`
    Name              string       `json:"name"`
    TenantID          string       `json:"tenant_id"`
    State             string       `json:"state"`
    ID                string       `json:"id"`
    RoutingTable      RoutingTable `json:"routingtable"`
    CreateTime        string       `json:"create_time"`
    AvailableIPCount  int          `json:"available_ip_count"`
    VPC               VPC          `json:"vpc"`
    VPCID             string       `json:"vpc_id"`
    Routes            []Route      `json:"routes"`
    Shared            bool         `json:"shared"`
    CIDR              string       `json:"cidr"`
    Gateway           string       `json:"gateway"`
}

type RoutingTable struct {
    GatewayID    string `json:"gateway_id"`
    DefaultTable bool   `json:"default_table"`
    Explicit     bool   `json:"explicit"`
    ID           string `json:"id"`
    Name         string `json:"name"`
}

type VPC struct {
    Shared bool   `json:"shared"`
    State  string `json:"state"`
    ID     string `json:"id"`
    CIDRv4 string `json:"cidrv4"`
    Name   string `json:"name"`
}

type Route struct {
    SubnetID string `json:"subnet_id"`
    TenantID string `json:"tenant_id"`
    Mask     int    `json:"mask"`
    Gateway  string `json:"gateway"`
    CIDR     string `json:"cidr"`
    ID       string `json:"id"`
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
		Links []gophercloud.Link `json:"vpcsubnets_links"`
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
		Subnets []Subnet `json:"vpcsubnets"` // Modified
	}
	err := (r.(SubnetPage)).ExtractInto(&s)
	return s.Subnets, err
}
