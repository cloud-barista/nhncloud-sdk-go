package testing

import (
	"fmt"
	"net/http"
	"testing"

	az "github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/blockstorage/extensions/availabilityzones"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper/client"
)

const GetOutput = `
{
    "availabilityZoneInfo": [
        {
            "zoneName": "internal",
            "zoneState": {
                "available": true
            }
        },
        {
            "zoneName": "nova",
            "zoneState": {
                "available": true
            }
        }
    ]
}`

var AZResult = []az.AvailabilityZone{
	{
		ZoneName:  "internal",
		ZoneState: az.ZoneState{Available: true},
	},
	{
		ZoneName:  "nova",
		ZoneState: az.ZoneState{Available: true},
	},
}

// HandleGetSuccessfully configures the test server to respond to a Get request
// for availability zone information.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-availability-zone", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}
