package v2

import (
	"github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/loadbalancer/v2/quotas"
)

var quotaUpdateOpts = quotas.UpdateOpts{
	Loadbalancer:  gophercloud.IntToPointer(25),
	Listener:      gophercloud.IntToPointer(45),
	Member:        gophercloud.IntToPointer(205),
	Pool:          gophercloud.IntToPointer(25),
	Healthmonitor: gophercloud.IntToPointer(5),
	L7Policy:      gophercloud.IntToPointer(55),
	L7Rule:        gophercloud.IntToPointer(105),
}
