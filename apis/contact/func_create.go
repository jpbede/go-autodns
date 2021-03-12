package contact

import (
	"context"
	"github.com/jpbede/go-autodns/apis/job"
	"github.com/jpbede/go-autodns/internal/transport"
)

func (c *client) Create(contact Contact, ctx context.Context) (*job.Response, error) {
	var out job.Response
	if err := c.transport.Post(ctx, "/contact", &out, transport.WithJSONRequestBody(contact)); err != nil {
		return nil, err
	}
	return &out, nil
}
