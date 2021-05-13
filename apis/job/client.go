package job

import (
	"go.bnck.me/autodns/internal/transport"
)

type client struct {
	transport *transport.Client
}

// New creates a new Client for the job Endpoint
func New(transport *transport.Client) Client {
	return &client{
		transport: transport,
	}
}
