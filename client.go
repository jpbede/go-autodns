package autodns

import (
	"github.com/jpbede/go-autodns/apis/contact"
	"github.com/jpbede/go-autodns/apis/domain"
	"github.com/jpbede/go-autodns/apis/job"
	"github.com/jpbede/go-autodns/internal/transport"
)

// Client represents the main client
type Client struct {
	transport *transport.Client

	contact contact.Client
	domain  domain.Client
	job     job.Client
}

// Contact returns a client related to contact operations
func (c *Client) Contact() contact.Client {
	if c.contact == nil {
		c.contact = contact.New(c.transport)
	}

	return c.contact
}

// Domain returns a client related to domain operations
func (c *Client) Domain() domain.Client {
	if c.domain == nil {
		c.domain = domain.New(c.transport)
	}

	return c.domain
}

// Job returns a client related to job operations
func (c *Client) Job() job.Client {
	if c.job == nil {
		c.job = job.New(c.transport)
	}

	return c.job
}
