package catalog

import (
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/identity/v3/tokens"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
)

// ServiceCatalogPage is a single page of Service results.
type ServiceCatalogPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the ServiceCatalogPage contains no results.
func (r ServiceCatalogPage) IsEmpty() (bool, error) {
	services, err := ExtractServiceCatalog(r)
	return len(services) == 0, err
}

// ExtractServiceCatalog extracts a slice of Catalog from a Collection acquired from List.
func ExtractServiceCatalog(r pagination.Page) ([]tokens.CatalogEntry, error) {
	var s struct {
		Entries []tokens.CatalogEntry `json:"catalog"`
	}
	err := (r.(ServiceCatalogPage)).ExtractInto(&s)
	return s.Entries, err
}
