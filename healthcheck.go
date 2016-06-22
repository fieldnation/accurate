package accurate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// HealthCheck represents a HTTP health check response.
type HealthCheck struct {
	Alive bool `json:"alive"`
}

// Alive checks if the accurate services is alive
func Alive() bool {

	// build url
	u, _ := url.Parse(URL.String())
	u.Path = path.Join(URL.Path, alive)

	// send request
	resp, err := http.Get(u.String())
	if err != nil {
		return false
	}

	// read the HTTP response body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	// unmarshal the health check
	var h HealthCheck
	if err = json.Unmarshal(b, &h); err != nil {
		return false
	}

	// check the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {
		return false
	}

	return h.Alive
}
