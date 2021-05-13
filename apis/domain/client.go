package domain

import (
	"go.bnck.me/autodns/internal/transport"
)

type client struct {
	transport *transport.Client
}

// New creates a new Client for the task Endpoint
func New(transport *transport.Client) Client {
	return &client{
		transport: transport,
	}
}
