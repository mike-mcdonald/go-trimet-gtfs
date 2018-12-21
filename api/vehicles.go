package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/mike-mcdonald/go-trimet-gtfs/types"
)

const serviceAddress = "https://developer.trimet.org/ws/v2/vehicles"

// RetrieveVehicles retrieves vehicle locations
// see https://developer.trimet.org/ws_docs/vehicle_locations_ws.shtml
func RetrieveVehicles(appID string, since int64) types.ResultSet {
	u, err := url.Parse(serviceAddress)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("appID", appID)
	if since != 0 {
		q.Set("since", string(since))
	}
	u.RawQuery = q.Encode()
	r, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	rs := types.ReturnValue{}
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &rs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return rs.ResultSet
}
