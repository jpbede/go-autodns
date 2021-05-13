package autodns

import (
	"github.com/rs/zerolog"
	"go.bnck.me/autodns/internal/transport"
	"net/http"
)

// ClientOption options for creating the client
type ClientOption func(*Client)

// WithAPIEndpoint supplies a API Endpoint to be used
func WithAPIEndpoint(api APIEndpoint) ClientOption {
	return func(c *Client) {
		c.transport.BaseURL = string(api)
	}
}

// WithHTTPClient supplies a already created http.Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.transport.HTTPClient = httpClient
	}
}

// WithLogger supplies a logger for the transport client
func WithLogger(logger *zerolog.Logger) ClientOption {
	return func(c *Client) {
		c.transport.SetLogger(logger)
	}
}

// WithCredentials supplies the api credentials
func WithCredentials(username, password string, context int) ClientOption {
	return func(c *Client) {
		c.transport.Credentials = &transport.APICredentials{
			Username: username,
			Password: password,
			Context:  context,
		}
	}
}
