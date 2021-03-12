package domain

import (
	"context"
)

func (c *client) CreateAuthInfo1(domain Domain, ctx context.Context) (*DomainResponse, error) {
	var out DomainResponse
	if err := c.transport.Post(ctx, "/domain/"+domain.Name+"/_authinfo1", &out); err != nil {
		return nil, err
	}
	return &out, nil
}
