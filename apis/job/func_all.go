package job

import (
	"context"
)

func (c *client) All(ctx context.Context) ([]*ObjectJob, error) {
	var out ObjectJobResponse
	if err := c.transport.Post(ctx, "/job/_search", &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}
