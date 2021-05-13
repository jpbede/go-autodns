package domain

import (
	"context"
	"go.bnck.me/autodns/apis/job"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) Update(domain Domain, ctx context.Context) (*job.Response, error) {
	var out job.Response
	if err := c.transport.Put(ctx, "/domain/"+domain.Name, &out, transport.WithJSONRequestBody(domain)); err != nil {
		return nil, err
	}
	return &out, nil
}
