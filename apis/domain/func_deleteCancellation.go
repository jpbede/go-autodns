package domain

import (
	"context"
	"github.com/jpbede/go-autodns/internal/transport"
)

func (c *client) DeleteCancellation(domain string, ctx context.Context) (*transport.BaseResponse, error) {
	var out transport.BaseResponse
	if err := c.transport.Delete(ctx, "/domain/"+domain+"/cancelation", &out); err != nil {
		return nil, err
	}
	return &out, nil
}
