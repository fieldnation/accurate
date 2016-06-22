package accurate

import (
	"net/url"
)

const (
	// defaults used to build the accurate API URL, these settings are
	// customizable and can be overridden via accurate.URL.
	scheme  = "https"
	host    = "api.accuratebackground.com"
	version = "v3"

	// endpoints used by different accurate API versions
	alive     = "/alive"
	candidate = "/candidate"
)

var (
	// clientID is the user to authenticate and authorize HTTP requests
	clientID string

	// clientSecret is the password to authenticate and authorize HTTP requests
	clientSecret string

	// URL is the url for the specific accurate API version chosen
	URL *url.URL
)

// the init method defaults to the latest accurate API version
func init() {
	URL = &url.URL{Scheme: scheme, Host: host, Path: version}
}

// SetClientID sets a client id.
func SetClientID(id string) {
	clientID = id
}

// SetClientSecret sets a client secret.
func SetClientSecret(secret string) {
	clientSecret = secret
}
