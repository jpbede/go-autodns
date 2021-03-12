package domain

import (
	"context"
)

func (c *client) All(ctx context.Context) ([]*Domain, error) {
	var out Response
	if err := c.transport.Post(ctx, "/domain/_search", &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}
