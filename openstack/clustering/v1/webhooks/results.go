package webhooks

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
)

type commonResult struct {
	gophercloud.Result
}

type TriggerResult struct {
	commonResult
}

// Extract retrieves the response action
func (r commonResult) Extract() (string, error) {
	var s struct {
		Action string `json:"action"`
	}
	err := r.ExtractInto(&s)
	return s.Action, err
}
