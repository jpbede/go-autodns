package autodns

import (
	"github.com/jpbede/go-autodns/apis/domain"
	"github.com/jpbede/go-autodns/internal/transport"
)

// Client represents the main client
type Client struct {
	transport *transport.Client

	domain domain.Client
}

// Stat returns a client related to stats
func (c *Client) Domain() domain.Client {
	if c.domain == nil {
		c.domain = domain.New(c.transport)
	}

	return c.domain
}
