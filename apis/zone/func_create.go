package zone

import (
	"context"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) Create(ctx context.Context, zone Zone) (*Response, error) {
	var out Response
	if err := c.transport.Post(ctx, "/zone", &out, transport.WithJSONRequestBody(zone)); err != nil {
		return nil, err
	}
	return &out, nil
}
