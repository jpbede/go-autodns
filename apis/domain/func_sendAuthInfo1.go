package domain

import (
	"context"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) SendAuthInfo1ToOwnerC(domain Domain, ctx context.Context) (*transport.BaseResponse, error) {
	var out transport.BaseResponse
	if err := c.transport.Post(ctx, "/domain/"+domain.Name+"/_sendAuthinfoToOwnerc", &out); err != nil {
		return nil, err
	}
	return &out, nil
}
