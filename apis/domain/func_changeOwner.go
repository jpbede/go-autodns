package domain

import (
	"context"
	"go.bnck.me/autodns/apis/contact"
	"go.bnck.me/autodns/apis/job"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) ChangeOwner(domain Domain, ownerC *contact.Contact, ctx context.Context) (*job.Response, error) {
	domain.OwnerC = ownerC // set OwnerC to new owner

	var out job.Response
	if err := c.transport.Put(ctx, "/domain/"+domain.Name+"/_ownerChange", &out, transport.WithJSONRequestBody(domain)); err != nil {
		return nil, err
	}
	return &out, nil
}
