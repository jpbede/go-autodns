package domain

import (
	"context"
	"github.com/jpbede/go-autodns/internal/transport"
)

func (c *client) DeleteAuthInfo1(domain Domain, ctx context.Context) (*transport.BaseResponse, error) {
	var out transport.BaseResponse
	if err := c.transport.Delete(ctx, "/domain/"+domain.Name+"/_authinfo1", &out); err != nil {
		return nil, err
	}
	return &out, nil
}
