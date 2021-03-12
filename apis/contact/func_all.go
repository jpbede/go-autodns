package contact

import "context"

func (c *client) All(ctx context.Context) ([]Contact, error) {
	var out Response
	if err := c.transport.Post(ctx, "/contact/_search", &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}
