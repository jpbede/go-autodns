package domain

import (
	"context"
)

func (c *client) Info(domain string, ctx context.Context) (*Domain, error) {
	var out DomainResponse
	if err := c.transport.Get(ctx, "/domain/"+domain, &out); err != nil {
		return nil, err
	}
	return &out.Data[0], nil
}
