package contact

import (
	"context"
	"github.com/jpbede/go-autodns/apis/job"
	"github.com/jpbede/go-autodns/internal/transport"
	"strconv"
)

func (c *client) UpdateComment(contact Contact, ctx context.Context) (*job.Response, error) {
	var out job.Response
	if err := c.transport.Put(ctx, "/contact/"+strconv.Itoa(contact.ID)+"/_comment", &out, transport.WithJSONRequestBody(contact)); err != nil {
		return nil, err
	}
	return &out, nil
}
