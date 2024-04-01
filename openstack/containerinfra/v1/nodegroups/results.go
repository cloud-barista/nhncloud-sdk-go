package nodegroups

import (
	"time"

	gophercloud "github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/pagination"
)

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*NodeGroup, error) {
	var s NodeGroup
	err := r.ExtractInto(&s)
	return &s, err
}

// GetResult is the response from a Get request.
// Use the Extract method to retrieve the NodeGroup itself.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create request.
// Use the Extract method to retrieve the created node group.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response from an Update request.
// Use the Extract method to retrieve the updated node group.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete request.
// Use the ExtractErr method to extract the error from the result.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpgradeResult is the response from an Upgrade request.
// Use the Extract method to retrieve the upgraded node group.
type UpgradeResult struct {
	commonResult
}

func (r UpgradeResult) Extract() (string, error) {
	var s struct {
		UUID string
	}
	err := r.ExtractInto(&s)
	return s.UUID, err
}

type Autoscale struct {
	CaEnable                 bool   `json:"ca_enable" required:"true"`
	CaImage                  string `json:"ca_image,omitempty"`
	CaMaxNodeCount           int    `json:"ca_max_node_count,omitempty"`
	CaMinNodeCount           int    `json:"ca_min_node_count,omitempty"`
	CaScaleDownEnable        bool   `json:"ca_scale_down_enable,omitempty"`
	CaScaleDownUnneededTime  int    `json:"ca_scale_down_unneeded_time,omitempty"`
	CaScaleDownUtilThresh    int    `json:"ca_scale_down_util_thresh,omitempty"`
	CaScaleDownDelayAfterAdd int    `json:"ca_scale_down_delay_after_add,omitempty"`
	Clusterautoscale         string `json:"clusterautoscale,omitempty"`
}

// GetAutoscaleResult is the response from a GetAutoscale request.
// Use the Extract method to extract from the result.
type GetAutoscaleResult struct {
	commonResult
}

// SetAutoscaleResult is the response from a SetAutoscale request.
// Use the Extract method to extract from the result.
type SetAutoscaleResult struct {
	commonResult
}

func (r GetAutoscaleResult) Extract() (*Autoscale, error) {
	var s Autoscale
	err := r.ExtractInto(&s)
	return &s, err
}

func (r SetAutoscaleResult) Extract() (string, error) {
	var s struct {
		UUID string
	}
	err := r.ExtractInto(&s)
	return s.UUID, err
}

// NodeGroup is the API representation of a Magnum node group.
type NodeGroup struct {
	ID               int                `json:"id"`
	UUID             string             `json:"uuid"`
	Name             string             `json:"name"`
	ClusterID        string             `json:"cluster_id"`
	ProjectID        string             `json:"project_id"`
	DockerVolumeSize *int               `json:"docker_volume_size"`
	Labels           map[string]string  `json:"labels"`
	Links            []gophercloud.Link `json:"links"`
	FlavorID         string             `json:"flavor_id"`
	ImageID          string             `json:"image_id"`
	NodeAddresses    []string           `json:"node_addresses"`
	NodeCount        int                `json:"node_count"`
	Role             string             `json:"role"`
	MinNodeCount     int                `json:"min_node_count"`
	MaxNodeCount     *int               `json:"max_node_count"`
	IsDefault        bool               `json:"is_default"`
	StackID          string             `json:"stack_id"`
	Status           string             `json:"status"`
	StatusReason     string             `json:"status_reason"`
	Version          string             `json:"version"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

type NodeGroupPage struct {
	pagination.LinkedPageBase
}

func (r NodeGroupPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, nil
}

func (r NodeGroupPage) IsEmpty() (bool, error) {
	s, err := ExtractNodeGroups(r)
	return len(s) == 0, err
}

// ExtractNodeGroups takes a Page of node groups as returned from List
// or from AllPages and extracts it as a slice of NodeGroups.
func ExtractNodeGroups(r pagination.Page) ([]NodeGroup, error) {
	var s struct {
		NodeGroups []NodeGroup `json:"nodegroups"`
	}
	err := (r.(NodeGroupPage)).ExtractInto(&s)
	return s.NodeGroups, err
}
