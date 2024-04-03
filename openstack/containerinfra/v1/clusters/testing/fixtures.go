package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	gophercloud "github.com/cloud-barista/nhncloud-sdk-go"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/containerinfra/v1/clusters"
	th "github.com/cloud-barista/nhncloud-sdk-go/testhelper"
	fake "github.com/cloud-barista/nhncloud-sdk-go/testhelper/client"
)

const clusterUUID = "746e779a-751a-456b-a3e9-c883d734946f"
const clusterUUID2 = "846e779a-751a-456b-a3e9-c883d734946f"
const requestUUID = "req-781e9bdc-4163-46eb-91c9-786c53188bbb"

var ClusterCreateResponse = fmt.Sprintf(`
										{
											"uuid":"%s"
										}`, clusterUUID)

var ExpectedCluster = clusters.Cluster{
	APIAddress:        "https://172.24.4.6:6443",
	COEVersion:        "v1.2.0",
	ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
	CreateTimeout:     60,
	CreatedAt:         time.Date(2016, 8, 29, 6, 51, 31, 0, time.UTC),
	DiscoveryURL:      "https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
	Links: []gophercloud.Link{
		{
			Href: "http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "self",
		},
		{
			Href: "http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "bookmark",
		},
	},
	KeyPair:            "my-keypair",
	MasterAddresses:    []string{"172.24.4.6"},
	MasterCount:        1,
	Name:               "k8s",
	NodeAddresses:      []string{"172.24.4.13"},
	NodeCount:          1,
	StackID:            "9c6f1169-7300-4d08-a444-d2be38758719",
	Status:             "CREATE_COMPLETE",
	StatusReason:       "Stack CREATE completed successfully",
	UpdatedAt:          time.Date(2016, 8, 29, 6, 53, 24, 0, time.UTC),
	UUID:               clusterUUID,
	FloatingIPEnabled:  true,
	FixedNetwork:       "private_network",
	FixedSubnet:        "private_subnet",
	HealthStatus:       "HEALTHY",
	HealthStatusReason: map[string]interface{}{"api": "ok"},
}

var ExpectedCluster2 = clusters.Cluster{
	APIAddress:        "https://172.24.4.6:6443",
	COEVersion:        "v1.2.0",
	ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
	CreateTimeout:     60,
	CreatedAt:         time.Time{},
	DiscoveryURL:      "https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
	Links: []gophercloud.Link{
		{
			Href: "http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "self",
		},
		{
			Href: "http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "bookmark",
		},
	},
	KeyPair:            "my-keypair",
	MasterAddresses:    []string{"172.24.4.6"},
	MasterCount:        1,
	Name:               "k8s",
	NodeAddresses:      []string{"172.24.4.13"},
	NodeCount:          1,
	StackID:            "9c6f1169-7300-4d08-a444-d2be38758719",
	Status:             "CREATE_COMPLETE",
	StatusReason:       "Stack CREATE completed successfully",
	UpdatedAt:          time.Date(2016, 8, 29, 6, 53, 24, 0, time.UTC),
	UUID:               clusterUUID2,
	FloatingIPEnabled:  true,
	FixedNetwork:       "private_network",
	FixedSubnet:        "private_subnet",
	HealthStatus:       "HEALTHY",
	HealthStatusReason: map[string]interface{}{"api": "ok"},
}

var ExpectedClusterUUID = clusterUUID

func HandleCreateClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprint(w, ClusterCreateResponse)
	})
}

func HandleGetClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ClusterGetResponse)
	})
}

var ClusterGetResponse = fmt.Sprintf(`
{
		"status":"CREATE_COMPLETE",
		"uuid":"%s",
		"links":[
		  {
			 "href":"http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			 "rel":"self"
		  },
		  {
			 "href":"http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			 "rel":"bookmark"
		  }
		],
		"stack_id":"9c6f1169-7300-4d08-a444-d2be38758719",
		"created_at":"2016-08-29T06:51:31+00:00",
		"api_address":"https://172.24.4.6:6443",
		"discovery_url":"https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
		"updated_at":"2016-08-29T06:53:24+00:00",
		"master_count":1,
		"coe_version": "v1.2.0",
		"keypair":"my-keypair",
		"cluster_template_id":"0562d357-8641-4759-8fed-8173f02c9633",
		"master_addresses":[
		  "172.24.4.6"
		],
		"node_count":1,
		"node_addresses":[
		  "172.24.4.13"
		],
		"status_reason":"Stack CREATE completed successfully",
		"create_timeout":60,
		"name":"k8s",
		"floating_ip_enabled": true,
		"fixed_network": "private_network",
		"fixed_subnet": "private_subnet",
		"health_status": "HEALTHY",
		"health_status_reason": {"api": "ok"}
}`, clusterUUID)

var ClusterListResponse = fmt.Sprintf(`
{
	"clusters": [
		{
			"api_address":"https://172.24.4.6:6443",
			"cluster_template_id":"0562d357-8641-4759-8fed-8173f02c9633",
			"coe_version": "v1.2.0",
			"create_timeout":60,
			"created_at":"2016-08-29T06:51:31+00:00",
			"discovery_url":"https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
			"keypair":"my-keypair",
			"links":[
			  {
				 "href":"http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"self"
			  },
			  {
				 "href":"http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"bookmark"
			  }
			],
			"master_addresses":[
			  "172.24.4.6"
			],
			"master_count":1,
			"name":"k8s",
			"node_addresses":[
			  "172.24.4.13"
			],
			"node_count":1,
			"stack_id":"9c6f1169-7300-4d08-a444-d2be38758719",
			"status":"CREATE_COMPLETE",
			"status_reason":"Stack CREATE completed successfully",
			"updated_at":"2016-08-29T06:53:24+00:00",
			"uuid":"%s",
			"floating_ip_enabled": true,
			"fixed_network": "private_network",
			"fixed_subnet": "private_subnet",
			"health_status": "HEALTHY",
			"health_status_reason": {"api": "ok"}
		},
		{
			"api_address":"https://172.24.4.6:6443",
			"cluster_template_id":"0562d357-8641-4759-8fed-8173f02c9633",
			"coe_version": "v1.2.0",
			"create_timeout":60,
			"created_at":null,
			"discovery_url":"https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
			"keypair":"my-keypair",
			"links":[
			  {
				 "href":"http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"self"
			  },
			  {
				 "href":"http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"bookmark"
			  }
			],
			"master_addresses":[
			  "172.24.4.6"
			],
			"master_count":1,
			"name":"k8s",
			"node_addresses":[
			  "172.24.4.13"
			],
			"node_count":1,
			"stack_id":"9c6f1169-7300-4d08-a444-d2be38758719",
			"status":"CREATE_COMPLETE",
			"status_reason":"Stack CREATE completed successfully",
			"updated_at":null,
			"uuid":"%s",
			"floating_ip_enabled": true,
			"fixed_network": "private_network",
			"fixed_subnet": "private_subnet",
			"health_status": "HEALTHY",
			"health_status_reason": {"api": "ok"}
		}
	]
}`, clusterUUID, clusterUUID2)

var ExpectedClusters = []clusters.Cluster{ExpectedCluster}

func HandleListClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ClusterListResponse)
	})
}

func HandleListDetailClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ClusterListResponse)
	})
}

var UpdateResponse = fmt.Sprintf(`
{
	"uuid":"%s"
}`, clusterUUID)

func HandleUpdateClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})
}

var UpgradeResponse = fmt.Sprintf(`
{
	"uuid":"%s"
}`, clusterUUID)

func HandleUpgradeClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID+"/actions/upgrade", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprint(w, UpgradeResponse)
	})
}

func HandleDeleteClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusNoContent)
	})
}

var ResizeResponse = fmt.Sprintf(`
{
	"uuid": "%s"
}`, clusterUUID)

func HandleResizeClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID+"/actions/resize", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprint(w, ResizeResponse)
	})
}

var ExpectedClusterConfig = "apiVersion: v1\nclusters:\n- cluster:\n    server: https://8.220.195.76:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURUakNDQWphZ0F3SUJBZ0lVT0llSWFxNDFqVlQrejRXU1pOZ1I3RFFvUXE0d0RRWUpLb1pJaHZjTkFRRUwKQlFBd1BqRW5NQThHQTFVRUNoTUlhR0Z1WjNwb2IzVXdGQVlEVlFRS0V3MWhiR2xpWVdKaElHTnNiM1ZrTVJNdwpFUVlEVlFRREV3cHJkV0psY201bGRHVnpNQ0FYRFRJME1ERXlPVEEyTkRNd01Gb1lEekl3TlRRd01USXhNRFkwCk16QXdXakErTVNjd0R3WURWUVFLRXdob1lXNW5lbWh2ZFRBVUJnTlZCQW9URFdGc2FXSmhZbUVnWTJ4dmRXUXgKRXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFDcWlBWXpMdzEyanhlRUtaNmNjTWxjcEJMV3h0N1AyTjNWRGdRakV6d0Y1MkRHN2ZVUGhBMm5EQWNvCncyK05FVXZUOG94MWZJUXg2VC9PK1psRkxzQXNxT3VweVNsNGV3dytFR3Y5RCtOYTlCV3VYaklZWm1FZEtVK0wKTzZWZ09rUTN2Y0F2M094N1lGcDNLa2ZRQXF3R0c1RUk1ZHBoRFlhY25obC9kLzR0YXRXUjEzc3FhTVI0N3QyUwpPYkhLclpKYmErWUc3V1MrbWlzbjczdzdwdm1LVHdIZ0JjVXBDTlFyZWNlKzNkWG5reS8ra2UzQys3VW1oRXFvClhjT2JyZUxlTko3dXpmb2hPT2ljamR6R05qdWFmYzI4dmtaNWJ5RDY4bmFiMWMvUTFjVFpYeUxXZ0Eyc3ZibEUKTEFqNWc1RGd1bWtkaXRETWl5WHBYeldBNnVkMUFnTUJBQUdqUWpCQU1BNEdBMVVkRHdFQi93UUVBd0lDckRBUApCZ05WSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUndFV2lQazQ3WUNlS2N6bTllOTVlTEVsTm53ekFOCkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQUN1NDlSU0ZlR2YrUmtUWlN2OWU4YlJrczZDSFFkRHErd05jRmVNYjgKWUJzdXZuZTdUMnNxb2ppcUZtak5haitJblBNeWFZOFRkUVpQSnl3VllZOEc4L0Jzd09MSlRQM1B6K3RUNzl0UQorL1ZERFMwYUU0MnhBSGhoNHdwTVg5alZtQXNtcVp6aXhJRjRISXhyU01CWDRvOGY5WnFudlo1UVE3cDZ4cEx1ClZmaDhQMUdvNlphZjViL3VCMGhlc3RyNndPdFRPbDM1T1ovM3dqVFU1RDRWTFFuK1U5QTVmcXlEa2FUc2hKU3cKemFUZUdBVUsvQno2UjZMazc1QzA4OEcrWGZiR3lVdDdRZXNCT1k1U2pjd3J4NzBYYjh5VzJ3UXJMSXMxVHNQYwordml3RFRBY0k4TFo0b1hXYk9MWXpTdFJWd1F6NEdPY0o0V3kxd0IvcDY4SGF3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"210335406175688605\"\n  name: 210335406175688605-c24a1916ee04f48ca96d910429b86a76f\ncurrent-context: 210335406175688605-c24a1916ee04f48ca96d910429b86a76f\nkind: Config\npreferences: {}\nusers:\n- name: \"210335406175688605\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQvekNDQXVlZ0F3SUJBZ0lFQVQvdlV6QU5CZ2txaGtpRzl3MEJBUXNGQURCcU1Tb3dLQVlEVlFRS0V5RmoKTWpSaE1Ua3hObVZsTURSbU5EaGpZVGsyWkRreE1EUXlPV0k0Tm1FM05tWXhFREFPQmdOVkJBc1RCMlJsWm1GMQpiSFF4S2pBb0JnTlZCQU1USVdNeU5HRXhPVEUyWldVd05HWTBPR05oT1Raa09URXdOREk1WWpnMllUYzJaakFlCkZ3MHlOREF4TWprd05qUXdNREJhRncweU56QXhNamd3TmpRMU5ESmFNRW94RlRBVEJnTlZCQW9UREhONWMzUmwKYlRwMWMyVnljekVKTUFjR0ExVUVDeE1BTVNZd0pBWURWUVFERXgweU1UQXpNelUwTURZeE56VTJPRGcyTURVdApNVGN3TmpVeE1EYzBNakNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNTVlDUG54CnM0Tm1YeVA1ZkJGZE9iNkw1bGxzeS81bFdIRnpHVzJkUi9rVHRZNW93YUJKaWVSMFVuS0o1ODBkUVFXVnVNeGUKb0NJRHlZQm1vaDZCaW1zZWljb1E3RkVoNG85amlubWYvRFJaSzdxckxsSTR2U1RicnBIaXN2ZkJCb1FTQkFTMwoxOHY0OVZ4dFVjaXlaRHlSVWpqckV4WHVRMnZkT3RQSVU2ekVPYVFxRHN2N3dLZE9TaXdNdUdGWWxHa29KSDNqCnMvZENQMjVDcXN5REtkK2s2VzcrZ0ZOeFZuQ3RmWkI1VENMelVBbjFGa3RCOENiK1V1OUhGRjRjQ0MzNUQ1N2YKRTRVdytrWWRDT0tyd3lpNnF1bXJrVml5M0xBRUtEWTJtOUV2UVQwTEZ0THVPR1VjMVRPZE50aDAxeDBIUnlGOQpyeHVHV3ltTForMmF5WFVDQXdFQUFhT0J6RENCeVRBT0JnTlZIUThCQWY4RUJBTUNCNEF3RXdZRFZSMGxCQXd3CkNnWUlLd1lCQlFVSEF3SXdEQVlEVlIwVEFRSC9CQUl3QURBZkJnTlZIU01FR0RBV2dCVGo5bkt0c2hZakNmUUIKa1BPbURSYUpQaHFUUlRBOEJnZ3JCZ0VGQlFjQkFRUXdNQzR3TEFZSUt3WUJCUVVITUFHR0lHaDBkSEE2THk5agpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnZiUzl2WTNOd01EVUdBMVVkSHdRdU1Dd3dLcUFvb0NhR0pHaDBkSEE2Ckx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnZiUzl5YjI5MExtTnliREFOQmdrcWhraUc5dzBCQVFzRkFBT0MKQVFFQWEzeTBvemhIdmxmMDU5TnY2c3R2Q1pOSG1MSXNhOG83U0Qrd0RSM3hPL2VFRGdvb1lHTjlycVY2T0tGTQpaejlpTzdGaDgwclZCOFAwVW8wWG1raGhCN0sxVFJnNlpNN0tSTWVWT1BuZmJwSWd3TU9wLzNXVjhGTTFFdm8zCjJrQUZuY3pETGJ3WDFoSHNJa1RlbkZBZUU0NDVna01zY0s0bXlsclRFTmpIZEIvYUpxa2NtUEwwVmZYRkthb2EKZUtYeUdJU1BQVEI3QkE2L3VQSFJRTEJycVgyd2FlTlRpQUpBOHRCazZqdkhoQnZXOURkb1BYUkVvVTVWekZpRAo5TFlPOWl2ZjFPU1MzTEFoR2xLVHFqckNoUVpHdnhoUmpLMzFEY1Nxakc2UkJ6Wkdyc05udGRwTFFOTHpYRXFFCmt0dy9CYTV0bmhWdHdqOVpXRFR1QUVTSGtnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQotLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJRG96Q0NBd3lnQXdJQkFnSUVBVC92VHpBTkJna3Foa2lHOXcwQkFRc0ZBREJpTVFzd0NRWURWUVFHRXdKRApUakVSTUE4R0ExVUVDQXdJV21obFNtbGhibWN4RVRBUEJnTlZCQWNNQ0VoaGJtZGFhRzkxTVJBd0RnWURWUVFLCkRBZEJiR2xpWVdKaE1Rd3dDZ1lEVlFRTERBTkJRMU14RFRBTEJnTlZCQU1NQkhKdmIzUXdIaGNOTWpRd01USTUKTURZME1EQXdXaGNOTkRRd01USTBNRFkwTlRNM1dqQnFNU293S0FZRFZRUUtFeUZqTWpSaE1Ua3hObVZsTURSbQpORGhqWVRrMlpEa3hNRFF5T1dJNE5tRTNObVl4RURBT0JnTlZCQXNUQjJSbFptRjFiSFF4S2pBb0JnTlZCQU1UCklXTXlOR0V4T1RFMlpXVXdOR1kwT0dOaE9UWmtPVEV3TkRJNVlqZzJZVGMyWmpDQ0FTSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1EaHZsRUg5T0czMG0xakdxQ0lnZnNsM2VJa1hBdkZrRTBkRzNQOApNWEV6UG1ZMmtYQTBJQXBBYlZTNHVmclFhSDlKZDZqODJjeGxZclN0WjNoSFc4N01IQVZENEFmQ3ZWVkZSdzd6CnRQUE5jQ1I3OFdJeDcyZDloaHFZeFl6Mit6NWxoQkZXS1MwTHdHQXI0SmZ5TnV1OHpkY1Y5czNGVVJBWW53Mk4KamszVEJjRlZJNXFTQkMvVm5acC90eXU2Wkd3dmdIQzBpWC9ybEUyVzY4blFaZzJjZHdsQm5oVEFlbDBWRWdYZwpESmM3aW4wZ2V3WlA3eEgxMkluT3hNbnV6N0lCSWJndC8yUExuK2IwcWhTVHNXVkwydFpQY2RJRG5WMytPUnhHCnVvQUt5VkZKdXlHZnYwbkxFUDZLTVowaU94OHVESTFpaVdlOS83QzczWlFMcVBFQ0F3RUFBYU9CMlRDQjFqQU8KQmdOVkhROEJBZjhFQkFNQ0Fxd3dEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVU0L1p5cmJJVwpJd24wQVpEenBnMFdpVDRhazBVd0h3WURWUjBqQkJnd0ZvQVVoVnIvM1NQTkpXOVlRVytlYlVVNW0xaDlkZjh3ClBBWUlLd1lCQlFVSEFRRUVNREF1TUN3R0NDc0dBUVVGQnpBQmhpQm9kSFJ3T2k4dlkyVnlkSE11WVdOekxtRnMKYVhsMWJpNWpiMjB2YjJOemNEQTFCZ05WSFI4RUxqQXNNQ3FnS0tBbWhpUm9kSFJ3T2k4dlkyVnlkSE11WVdOegpMbUZzYVhsMWJpNWpiMjB2Y205dmRDNWpjbXd3RFFZSktvWklodmNOQVFFTEJRQURnWUVBSHEwSGY2cUZ1Unl3CisyRGpiS2xYVUkzV2tlVHdva1ZWVnd6OGVMd0VYSm52MFZQTjRndTlUR1N4UG1saFJaRHpNaStGWnpvUUMrdUMKQ1lqYkhHdmM3bGNUNXQ5bHgxbThpbnpYVks2S3Zxa29RUCtYeWg4ZFRjeVRVazBnTjZ2djcxOVc5OHFBWjJoVgp2Vm82M1UrV1owLytpZVV0QytPaWV6VW13dFFXY0RJPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcGdJQkFBS0NBUUVBd3hnSStmR3pnMlpmSS9sOEVWMDV2b3ZtV1d6TC9tVlljWE1aYloxSCtSTzFqbWpCCm9FbUo1SFJTY29ubnpSMUJCWlc0ekY2Z0lnUEpnR2FpSG9HS2F4Nkp5aERzVVNIaWoyT0tlWi84TkZrcnVxc3UKVWppOUpOdXVrZUt5OThFR2hCSUVCTGZYeS9qMVhHMVJ5TEprUEpGU09Pc1RGZTVEYTkwNjA4aFRyTVE1cENvTwp5L3ZBcDA1S0xBeTRZVmlVYVNna2ZlT3o5MEkvYmtLcXpJTXAzNlRwYnY2QVUzRldjSzE5a0hsTUl2TlFDZlVXClMwSHdKdjVTNzBjVVhod0lMZmtQbnQ4VGhURDZSaDBJNHF2REtMcXE2YXVSV0xMY3NBUW9OamFiMFM5QlBRc1cKMHU0NFpSelZNNTAyMkhUWEhRZEhJWDJ2RzRaYktZdG43WnJKZFFJREFRQUJBb0lCQVFDbHNvNDlLd25hZW1JbgovY2RnUUJ2Qk9MVW1FbitYeTUrNGk1ZDNQL1JYR21SRFZibk9WMmNrZjU5ZTVMM1p1aFQxbDFwalNhTWNBTGR0CnRMM1F4eGszTGJLOXltM0w0c200RzBVTFU5UXk2ZUY0STlldHdiN3p4Um1aa2orS3FZak5OemZ1Z2U1b3NVUkQKZ2gzS0phbDR2N3hqMTFSWWRnc1JXODZRL1Q0RUp3K0tyenlaM3pNMXFmclg2UkdJRDQ2dUxZM3VlZS9wYVFyMApUa29Pc2dSUU5lZTZ3bFhVTGN6UkRFTUh4TXhrWnhVdnNyVWYzNVQ3Y21HaUxDMDFRTmJDZWg4ZStiQ0huQmVNCkFId1hBelczM2lQYWFrRFFFVytVN3VPMXlFNXErZGl1Qm9UcmtDUTFmMisveTdySklBbjJmZmQ3bHNJTTRWVGQKeGtIVHcvWUJBb0dCQVBhNWpUZEdwalNyMkxJTnNvYnVKVWM5M0ROejlDRERCRUZrMG9LcDNTTkttK05jSCs0WQo4RUkwMFB2ZkgrT2U3WXVkaWR0Y01PbEo5Z0dlSHBsd2kzbUJWY3NDclpMaXp5VWZiZU9SMEtrS1JTT2dxbitCCkQvbkh0Tk9EeDQ3UkFUMnM5ZUNmVkUvS201RUcvTFBXNnNmZ29wWGo2M0JPYkpPeGw2RHFub2cxQW9HQkFNcHQKbUEyRTk1YnhXRWt0dVNRd29ISXFqTzExNDNvMHRQWmpCa3ArOW8ySHpwWXpNNCsrRzJuOGhERlJOV2V6WXl0aQo5RWhBSERMclBOcVNtS09LbTd1NDBsSE9CR29Ra1ErYk9qbUxua25IcnFrdHV3SmNKL2hXaGN6NHpTRi9lK3FpCnhMaTU1SUlzRFc5RFZ4WnZBN3NLYlpaWXRKTkgzWFVYanF1eFArUkJBb0dCQUpwYjJGZmU0WmxPY2xKamMzQXIKaWpNYUpxd2lQWGhKeWwzSFlGVFRSVUVSS3BxQ1JvL0dGbnExWkpKUU5EbEtjei9JSGptWlloaHlaM0QzcGhsRgphbDFvWjI2TmpGNjdlL0d6eWlKNFZkMk1TVmxTNlppLy9HaS9Zd2g2QjkzNE1SaVBIMzJhOFRyQ2ZiV3NjMmxvCmFwY0dtWEhCbU1rNHA0RjN1Rys2bWRpSkFvR0JBSWlCTXdJYTRtTTFTTlhBTVpOSC8wbzlpMTh1R2tIMGZQdVIKZDFLUkgyMzlZTUJFc3NhQjZqYWtnL2hGelArckpuckJkZjJxemRsQWJIN3dVR0lvUERCd3g5TkdYQSt5TUVBWAo2MEdXOXh6RnZQQjAyQ0VWU0JiZE1ja3hGaE02eHJOSkkrTjE1N1FvUkw3ODZDTkdwWHRoNHRoNTVQdFVnUEJPCi9IU25TSENCQW9HQkFJZXF3Nmw0YzZtZkRVUU5GR2JPck1EQUZYdmxyMEVQdUp0bVRPcjBDaEIzMWpjWFFTQ2kKb3JLYWZIcE9DTG5KUW5QWndJQ2UzNlRoNGVYcVJQVy8wYi9TOURrbVNTdzRnSTNvWTJuQU1MUzNsQ2h2N0RadwpRR2p0cjd5dCt0TFdjYjhza3AvVmI5Q0pVSkpMZzdaVVdKODlKVEJ3QVRwaUNCZHk3SzNyUTJ0eQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

var GetConfigResponse = fmt.Sprintf(`
{
	"config":"apiVersion: v1\nclusters:\n- cluster:\n    server: https://8.220.195.76:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURUakNDQWphZ0F3SUJBZ0lVT0llSWFxNDFqVlQrejRXU1pOZ1I3RFFvUXE0d0RRWUpLb1pJaHZjTkFRRUwKQlFBd1BqRW5NQThHQTFVRUNoTUlhR0Z1WjNwb2IzVXdGQVlEVlFRS0V3MWhiR2xpWVdKaElHTnNiM1ZrTVJNdwpFUVlEVlFRREV3cHJkV0psY201bGRHVnpNQ0FYRFRJME1ERXlPVEEyTkRNd01Gb1lEekl3TlRRd01USXhNRFkwCk16QXdXakErTVNjd0R3WURWUVFLRXdob1lXNW5lbWh2ZFRBVUJnTlZCQW9URFdGc2FXSmhZbUVnWTJ4dmRXUXgKRXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFDcWlBWXpMdzEyanhlRUtaNmNjTWxjcEJMV3h0N1AyTjNWRGdRakV6d0Y1MkRHN2ZVUGhBMm5EQWNvCncyK05FVXZUOG94MWZJUXg2VC9PK1psRkxzQXNxT3VweVNsNGV3dytFR3Y5RCtOYTlCV3VYaklZWm1FZEtVK0wKTzZWZ09rUTN2Y0F2M094N1lGcDNLa2ZRQXF3R0c1RUk1ZHBoRFlhY25obC9kLzR0YXRXUjEzc3FhTVI0N3QyUwpPYkhLclpKYmErWUc3V1MrbWlzbjczdzdwdm1LVHdIZ0JjVXBDTlFyZWNlKzNkWG5reS8ra2UzQys3VW1oRXFvClhjT2JyZUxlTko3dXpmb2hPT2ljamR6R05qdWFmYzI4dmtaNWJ5RDY4bmFiMWMvUTFjVFpYeUxXZ0Eyc3ZibEUKTEFqNWc1RGd1bWtkaXRETWl5WHBYeldBNnVkMUFnTUJBQUdqUWpCQU1BNEdBMVVkRHdFQi93UUVBd0lDckRBUApCZ05WSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUndFV2lQazQ3WUNlS2N6bTllOTVlTEVsTm53ekFOCkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQUN1NDlSU0ZlR2YrUmtUWlN2OWU4YlJrczZDSFFkRHErd05jRmVNYjgKWUJzdXZuZTdUMnNxb2ppcUZtak5haitJblBNeWFZOFRkUVpQSnl3VllZOEc4L0Jzd09MSlRQM1B6K3RUNzl0UQorL1ZERFMwYUU0MnhBSGhoNHdwTVg5alZtQXNtcVp6aXhJRjRISXhyU01CWDRvOGY5WnFudlo1UVE3cDZ4cEx1ClZmaDhQMUdvNlphZjViL3VCMGhlc3RyNndPdFRPbDM1T1ovM3dqVFU1RDRWTFFuK1U5QTVmcXlEa2FUc2hKU3cKemFUZUdBVUsvQno2UjZMazc1QzA4OEcrWGZiR3lVdDdRZXNCT1k1U2pjd3J4NzBYYjh5VzJ3UXJMSXMxVHNQYwordml3RFRBY0k4TFo0b1hXYk9MWXpTdFJWd1F6NEdPY0o0V3kxd0IvcDY4SGF3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"210335406175688605\"\n  name: 210335406175688605-c24a1916ee04f48ca96d910429b86a76f\ncurrent-context: 210335406175688605-c24a1916ee04f48ca96d910429b86a76f\nkind: Config\npreferences: {}\nusers:\n- name: \"210335406175688605\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQvekNDQXVlZ0F3SUJBZ0lFQVQvdlV6QU5CZ2txaGtpRzl3MEJBUXNGQURCcU1Tb3dLQVlEVlFRS0V5RmoKTWpSaE1Ua3hObVZsTURSbU5EaGpZVGsyWkRreE1EUXlPV0k0Tm1FM05tWXhFREFPQmdOVkJBc1RCMlJsWm1GMQpiSFF4S2pBb0JnTlZCQU1USVdNeU5HRXhPVEUyWldVd05HWTBPR05oT1Raa09URXdOREk1WWpnMllUYzJaakFlCkZ3MHlOREF4TWprd05qUXdNREJhRncweU56QXhNamd3TmpRMU5ESmFNRW94RlRBVEJnTlZCQW9UREhONWMzUmwKYlRwMWMyVnljekVKTUFjR0ExVUVDeE1BTVNZd0pBWURWUVFERXgweU1UQXpNelUwTURZeE56VTJPRGcyTURVdApNVGN3TmpVeE1EYzBNakNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNTVlDUG54CnM0Tm1YeVA1ZkJGZE9iNkw1bGxzeS81bFdIRnpHVzJkUi9rVHRZNW93YUJKaWVSMFVuS0o1ODBkUVFXVnVNeGUKb0NJRHlZQm1vaDZCaW1zZWljb1E3RkVoNG85amlubWYvRFJaSzdxckxsSTR2U1RicnBIaXN2ZkJCb1FTQkFTMwoxOHY0OVZ4dFVjaXlaRHlSVWpqckV4WHVRMnZkT3RQSVU2ekVPYVFxRHN2N3dLZE9TaXdNdUdGWWxHa29KSDNqCnMvZENQMjVDcXN5REtkK2s2VzcrZ0ZOeFZuQ3RmWkI1VENMelVBbjFGa3RCOENiK1V1OUhGRjRjQ0MzNUQ1N2YKRTRVdytrWWRDT0tyd3lpNnF1bXJrVml5M0xBRUtEWTJtOUV2UVQwTEZ0THVPR1VjMVRPZE50aDAxeDBIUnlGOQpyeHVHV3ltTForMmF5WFVDQXdFQUFhT0J6RENCeVRBT0JnTlZIUThCQWY4RUJBTUNCNEF3RXdZRFZSMGxCQXd3CkNnWUlLd1lCQlFVSEF3SXdEQVlEVlIwVEFRSC9CQUl3QURBZkJnTlZIU01FR0RBV2dCVGo5bkt0c2hZakNmUUIKa1BPbURSYUpQaHFUUlRBOEJnZ3JCZ0VGQlFjQkFRUXdNQzR3TEFZSUt3WUJCUVVITUFHR0lHaDBkSEE2THk5agpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnZiUzl2WTNOd01EVUdBMVVkSHdRdU1Dd3dLcUFvb0NhR0pHaDBkSEE2Ckx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnZiUzl5YjI5MExtTnliREFOQmdrcWhraUc5dzBCQVFzRkFBT0MKQVFFQWEzeTBvemhIdmxmMDU5TnY2c3R2Q1pOSG1MSXNhOG83U0Qrd0RSM3hPL2VFRGdvb1lHTjlycVY2T0tGTQpaejlpTzdGaDgwclZCOFAwVW8wWG1raGhCN0sxVFJnNlpNN0tSTWVWT1BuZmJwSWd3TU9wLzNXVjhGTTFFdm8zCjJrQUZuY3pETGJ3WDFoSHNJa1RlbkZBZUU0NDVna01zY0s0bXlsclRFTmpIZEIvYUpxa2NtUEwwVmZYRkthb2EKZUtYeUdJU1BQVEI3QkE2L3VQSFJRTEJycVgyd2FlTlRpQUpBOHRCazZqdkhoQnZXOURkb1BYUkVvVTVWekZpRAo5TFlPOWl2ZjFPU1MzTEFoR2xLVHFqckNoUVpHdnhoUmpLMzFEY1Nxakc2UkJ6Wkdyc05udGRwTFFOTHpYRXFFCmt0dy9CYTV0bmhWdHdqOVpXRFR1QUVTSGtnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQotLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJRG96Q0NBd3lnQXdJQkFnSUVBVC92VHpBTkJna3Foa2lHOXcwQkFRc0ZBREJpTVFzd0NRWURWUVFHRXdKRApUakVSTUE4R0ExVUVDQXdJV21obFNtbGhibWN4RVRBUEJnTlZCQWNNQ0VoaGJtZGFhRzkxTVJBd0RnWURWUVFLCkRBZEJiR2xpWVdKaE1Rd3dDZ1lEVlFRTERBTkJRMU14RFRBTEJnTlZCQU1NQkhKdmIzUXdIaGNOTWpRd01USTUKTURZME1EQXdXaGNOTkRRd01USTBNRFkwTlRNM1dqQnFNU293S0FZRFZRUUtFeUZqTWpSaE1Ua3hObVZsTURSbQpORGhqWVRrMlpEa3hNRFF5T1dJNE5tRTNObVl4RURBT0JnTlZCQXNUQjJSbFptRjFiSFF4S2pBb0JnTlZCQU1UCklXTXlOR0V4T1RFMlpXVXdOR1kwT0dOaE9UWmtPVEV3TkRJNVlqZzJZVGMyWmpDQ0FTSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1EaHZsRUg5T0czMG0xakdxQ0lnZnNsM2VJa1hBdkZrRTBkRzNQOApNWEV6UG1ZMmtYQTBJQXBBYlZTNHVmclFhSDlKZDZqODJjeGxZclN0WjNoSFc4N01IQVZENEFmQ3ZWVkZSdzd6CnRQUE5jQ1I3OFdJeDcyZDloaHFZeFl6Mit6NWxoQkZXS1MwTHdHQXI0SmZ5TnV1OHpkY1Y5czNGVVJBWW53Mk4KamszVEJjRlZJNXFTQkMvVm5acC90eXU2Wkd3dmdIQzBpWC9ybEUyVzY4blFaZzJjZHdsQm5oVEFlbDBWRWdYZwpESmM3aW4wZ2V3WlA3eEgxMkluT3hNbnV6N0lCSWJndC8yUExuK2IwcWhTVHNXVkwydFpQY2RJRG5WMytPUnhHCnVvQUt5VkZKdXlHZnYwbkxFUDZLTVowaU94OHVESTFpaVdlOS83QzczWlFMcVBFQ0F3RUFBYU9CMlRDQjFqQU8KQmdOVkhROEJBZjhFQkFNQ0Fxd3dEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVU0L1p5cmJJVwpJd24wQVpEenBnMFdpVDRhazBVd0h3WURWUjBqQkJnd0ZvQVVoVnIvM1NQTkpXOVlRVytlYlVVNW0xaDlkZjh3ClBBWUlLd1lCQlFVSEFRRUVNREF1TUN3R0NDc0dBUVVGQnpBQmhpQm9kSFJ3T2k4dlkyVnlkSE11WVdOekxtRnMKYVhsMWJpNWpiMjB2YjJOemNEQTFCZ05WSFI4RUxqQXNNQ3FnS0tBbWhpUm9kSFJ3T2k4dlkyVnlkSE11WVdOegpMbUZzYVhsMWJpNWpiMjB2Y205dmRDNWpjbXd3RFFZSktvWklodmNOQVFFTEJRQURnWUVBSHEwSGY2cUZ1Unl3CisyRGpiS2xYVUkzV2tlVHdva1ZWVnd6OGVMd0VYSm52MFZQTjRndTlUR1N4UG1saFJaRHpNaStGWnpvUUMrdUMKQ1lqYkhHdmM3bGNUNXQ5bHgxbThpbnpYVks2S3Zxa29RUCtYeWg4ZFRjeVRVazBnTjZ2djcxOVc5OHFBWjJoVgp2Vm82M1UrV1owLytpZVV0QytPaWV6VW13dFFXY0RJPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcGdJQkFBS0NBUUVBd3hnSStmR3pnMlpmSS9sOEVWMDV2b3ZtV1d6TC9tVlljWE1aYloxSCtSTzFqbWpCCm9FbUo1SFJTY29ubnpSMUJCWlc0ekY2Z0lnUEpnR2FpSG9HS2F4Nkp5aERzVVNIaWoyT0tlWi84TkZrcnVxc3UKVWppOUpOdXVrZUt5OThFR2hCSUVCTGZYeS9qMVhHMVJ5TEprUEpGU09Pc1RGZTVEYTkwNjA4aFRyTVE1cENvTwp5L3ZBcDA1S0xBeTRZVmlVYVNna2ZlT3o5MEkvYmtLcXpJTXAzNlRwYnY2QVUzRldjSzE5a0hsTUl2TlFDZlVXClMwSHdKdjVTNzBjVVhod0lMZmtQbnQ4VGhURDZSaDBJNHF2REtMcXE2YXVSV0xMY3NBUW9OamFiMFM5QlBRc1cKMHU0NFpSelZNNTAyMkhUWEhRZEhJWDJ2RzRaYktZdG43WnJKZFFJREFRQUJBb0lCQVFDbHNvNDlLd25hZW1JbgovY2RnUUJ2Qk9MVW1FbitYeTUrNGk1ZDNQL1JYR21SRFZibk9WMmNrZjU5ZTVMM1p1aFQxbDFwalNhTWNBTGR0CnRMM1F4eGszTGJLOXltM0w0c200RzBVTFU5UXk2ZUY0STlldHdiN3p4Um1aa2orS3FZak5OemZ1Z2U1b3NVUkQKZ2gzS0phbDR2N3hqMTFSWWRnc1JXODZRL1Q0RUp3K0tyenlaM3pNMXFmclg2UkdJRDQ2dUxZM3VlZS9wYVFyMApUa29Pc2dSUU5lZTZ3bFhVTGN6UkRFTUh4TXhrWnhVdnNyVWYzNVQ3Y21HaUxDMDFRTmJDZWg4ZStiQ0huQmVNCkFId1hBelczM2lQYWFrRFFFVytVN3VPMXlFNXErZGl1Qm9UcmtDUTFmMisveTdySklBbjJmZmQ3bHNJTTRWVGQKeGtIVHcvWUJBb0dCQVBhNWpUZEdwalNyMkxJTnNvYnVKVWM5M0ROejlDRERCRUZrMG9LcDNTTkttK05jSCs0WQo4RUkwMFB2ZkgrT2U3WXVkaWR0Y01PbEo5Z0dlSHBsd2kzbUJWY3NDclpMaXp5VWZiZU9SMEtrS1JTT2dxbitCCkQvbkh0Tk9EeDQ3UkFUMnM5ZUNmVkUvS201RUcvTFBXNnNmZ29wWGo2M0JPYkpPeGw2RHFub2cxQW9HQkFNcHQKbUEyRTk1YnhXRWt0dVNRd29ISXFqTzExNDNvMHRQWmpCa3ArOW8ySHpwWXpNNCsrRzJuOGhERlJOV2V6WXl0aQo5RWhBSERMclBOcVNtS09LbTd1NDBsSE9CR29Ra1ErYk9qbUxua25IcnFrdHV3SmNKL2hXaGN6NHpTRi9lK3FpCnhMaTU1SUlzRFc5RFZ4WnZBN3NLYlpaWXRKTkgzWFVYanF1eFArUkJBb0dCQUpwYjJGZmU0WmxPY2xKamMzQXIKaWpNYUpxd2lQWGhKeWwzSFlGVFRSVUVSS3BxQ1JvL0dGbnExWkpKUU5EbEtjei9JSGptWlloaHlaM0QzcGhsRgphbDFvWjI2TmpGNjdlL0d6eWlKNFZkMk1TVmxTNlppLy9HaS9Zd2g2QjkzNE1SaVBIMzJhOFRyQ2ZiV3NjMmxvCmFwY0dtWEhCbU1rNHA0RjN1Rys2bWRpSkFvR0JBSWlCTXdJYTRtTTFTTlhBTVpOSC8wbzlpMTh1R2tIMGZQdVIKZDFLUkgyMzlZTUJFc3NhQjZqYWtnL2hGelArckpuckJkZjJxemRsQWJIN3dVR0lvUERCd3g5TkdYQSt5TUVBWAo2MEdXOXh6RnZQQjAyQ0VWU0JiZE1ja3hGaE02eHJOSkkrTjE1N1FvUkw3ODZDTkdwWHRoNHRoNTVQdFVnUEJPCi9IU25TSENCQW9HQkFJZXF3Nmw0YzZtZkRVUU5GR2JPck1EQUZYdmxyMEVQdUp0bVRPcjBDaEIzMWpjWFFTQ2kKb3JLYWZIcE9DTG5KUW5QWndJQ2UzNlRoNGVYcVJQVy8wYi9TOURrbVNTdzRnSTNvWTJuQU1MUzNsQ2h2N0RadwpRR2p0cjd5dCt0TFdjYjhza3AvVmI5Q0pVSkpMZzdaVVdKODlKVEJ3QVRwaUNCZHk3SzNyUTJ0eQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
}`)

func HandleGetConfigClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID+"/config", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetConfigResponse)
	})
}
