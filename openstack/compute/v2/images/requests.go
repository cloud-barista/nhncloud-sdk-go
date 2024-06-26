package images

import (
	"strings"

	gophercloud "github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// ListDetail request.
type ListOptsBuilder interface {
	ToImageListQuery() (string, error)
}

// ListOpts contain options filtering Images returned from a call to ListDetail.
type ListOpts struct {
	// ChangesSince filters Images based on the last changed status (in date-time
	// format).
	ChangesSince string `q:"changes-since"`

	// Limit limits the number of Images to return.
	Limit int `q:"limit"`

	// Mark is an Image UUID at which to set a marker.
	Marker string `q:"marker"`

	// Name is the name of the Image.
	Name string `q:"name"`

	// Server is the name of the Server (in URL format).
	Server string `q:"server"`

	// Status is the current status of the Image.
	Status string `q:"status"`

	// Type is the type of image (e.g. BASE, SERVER, ALL).
	Type string `q:"type"`
}

// ToImageListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToImageListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the available images. // $$$ Added
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {

	urlBefore := listURL(client)

	//Check!!
	// log.Println("\n### URL in ListDetail : ", urlBefore)

	// url := strings.Replace(urlBefore, "/00992ba6139a4fa4a3786701a6947525", "", -1)
	// url := strings.Replace(urlBefore, "/detail", "", -1)

	// $$$
	url := strings.Replace(urlBefore, "/images/", "/images", -1)

	//Check!!
	//log.Println("\n### List Query URL : ", url)

	if opts != nil {
		query, err := opts.ToImageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ImagePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ListDetail enumerates the available images.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {

	url := listDetailURL(client)

	//To Check!!
	//log.Println("\n### List Detail Query URL : ", url) // Added

	if opts != nil {
		query, err := opts.ToImageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ImagePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get returns data about a specific image by its ID.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {

	//Check!!
	url := getURL(client, id) // Modified
	//log.Println("\n### Get Query URL : ", url)

	// resp, err := client.Get(getURL(client, id), &r.Body, nil)
	resp, err := client.Get(url, &r.Body, nil) // Modified
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes the specified image ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
