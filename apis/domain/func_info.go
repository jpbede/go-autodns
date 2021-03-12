package domain

import (
	"context"
	"errors"
)

// Info gets information for a specific domain
func (c *client) Info(domain string, ctx context.Context) (*Domain, error) {
	var out Response
	if err := c.transport.Get(ctx, "/domain/"+domain, &out); err != nil {
		return nil, err
	}
	if len(out.Data) > 0 {
		return out.Data[0], nil
	} else {
		return nil, errors.New("no domain found")
	}
}
