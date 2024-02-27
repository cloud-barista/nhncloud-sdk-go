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
	"encoding/json"
	"time"

	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a network resource.
func (r commonResult) Extract() (*VPC, error) {
	var s VPC
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "vpc")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a VPC.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a VPC.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a VPC.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type RoutingTable struct {
    GatewayID    	string    `json:"gateway_id"`
    Subnets      	[]Subnet  `json:"subnets"`
    Name         	string    `json:"name"`
    VPCs         	[]VPC     `json:"vpcs"`
    TenantID     	string    `json:"tenant_id"`
    Distributed  	bool      `json:"distributed"`
    State        	string    `json:"state"`
    DefaultTable 	bool      `json:"default_table"`
    CreateTime   	time.Time `json:"create_time"`
    Routes       	[]Route   `json:"routes"`
    ID           	string    `json:"id"`
}

type Route struct {
	SubnetID 		string `json:"subnet_id"`
	TenantID 		string `json:"tenant_id"`
	Mask     		int    `json:"mask"`
	Gateway  		string `json:"gateway"`
	GatewayID      	string `json:"gateway_id,omitempty"`
	RoutingTableID 	string `json:"routingtable_id"`
	CIDR     		string `json:"cidr"`
	ID       		string `json:"id"`
}

type Subnet struct {
	RouterExternal 	bool    `json:"router:external"`
	Name           	string  `json:"name"`
	EnableDHCP     	bool    `json:"enable_dhcp"`
	TenantID       	string  `json:"tenant_id"`
	Gateway        	string  `json:"gateway"`
	Routes         	[]Route `json:"routes"`
	State          	string  `json:"state"`
	CreateTime     	string  `json:"create_time"`
	AvailableIPCount int   `json:"available_ip_count"`
	VPC            	VPC     `json:"vpc"`
	Shared         	bool    `json:"shared"`
	ID             	string  `json:"id"`
	VPCID          	string  `json:"vpc_id"`
	Hidden         	bool    `json:"hidden"`
	CIDR           	string  `json:"cidr"`
}

type VPC struct {
	RouterExternal 	bool     		`json:"router:external"`
	RoutingTables  	[]RoutingTable 	`json:"routingtables"`
	Name           	string   		`json:"name"`
	Subnets        	[]Subnet 		`json:"subnets"`
	TenantID       	string   		`json:"tenant_id"`
	State          	string   		`json:"state"`
	CreateTime     	string   		`json:"create_time"`
	CIDRv4         	string   		`json:"cidrv4"`
	Shared         	bool     		`json:"shared"`
	ID             	string   		`json:"id"`
}

func (r *VPC) UnmarshalJSON(b []byte) error {
	type tmp VPC

	// Support for older neutron time format
	var s struct {
		tmp
	}

	err := json.Unmarshal(b, &s)
	if err == nil {
		*r = VPC(s.tmp)
		return nil
	}

	return err // Caution!!
}

// VPCPage is the page returned by a pager when traversing over a
// collection of networks.
type VPCPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r VPCPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"vpcs_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a VPCPage struct is empty.
func (r VPCPage) IsEmpty() (bool, error) {
	is, err := ExtractVPCs(r)
	return len(is) == 0, err
}

// ExtractVPCs accepts a Page struct, specifically a VPCPage struct,
// and extracts the elements into a slice of VPC structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractVPCs(r pagination.Page) ([]VPC, error) {
	var s []VPC
	err := ExtractVPCsInto(r, &s)
	return s, err
}

func ExtractVPCsInto(r pagination.Page, v interface{}) error {
	return r.(VPCPage).Result.ExtractIntoSlicePtr(v, "vpcs")
}
