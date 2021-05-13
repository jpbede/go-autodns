package autodns

import (
	"errors"
	"go.bnck.me/autodns/internal/transport"
	"net/http"
)

// APIEndpoint represents a api endpoint
type APIEndpoint string

const (
	// APIEndpointProduction Production API endpoint
	APIEndpointProduction APIEndpoint = "https://api.autodns.com/v1"
	// APIEndpointDemo Demo API endpoint
	APIEndpointDemo APIEndpoint = "https://api.demo.autodns.com/v1"
)

// New creates a new Client with username, password and context
func New(username, password string, context int) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithCredentials(username, password, context))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client, username, password string, context int) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithHTTPClient(httpClient), WithCredentials(username, password, context))
}

// NewWithOptions creates a new Client with given options
func NewWithOptions(options ...ClientOption) (*Client, error) {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.New(string(APIEndpointProduction))

	// run given options
	for _, option := range options {
		option(c)
	}

	// check if there are credentials
	if !c.transport.HasCredentials() {
		return nil, errors.New("no api credentials supplied")
	}

	return c, nil
}
