package zone

import "context"

func (c *client) All(ctx context.Context) ([]*Zone, error) {
	var out Response
	if err := c.transport.Post(ctx, "/zone/_search", &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}
