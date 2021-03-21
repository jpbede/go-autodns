package domain

import (
	"context"
	"errors"
	"github.com/jpbede/go-autodns/internal/transport"
)

func (c *client) Cancel(domain Domain, ctx context.Context) (*Cancellation, error) {
	var out *CancellationResponse
	if err := c.transport.Post(ctx, "/domain/"+domain.Name+"/cancelation", &out); err != nil {
		return nil, err
	}
	if len(out.Data) > 0 {
		return out.Data[0], nil
	}
	return nil, errors.New("no cancellation found")
}

func (c *client) GetCancellations(ctx context.Context) ([]*Cancellation, error) {
	var out CancellationResponse
	if err := c.transport.Post(ctx, "/domain/cancelation/_search", &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

func (c *client) DeleteCancellation(domain string, ctx context.Context) (*transport.BaseResponse, error) {
	var out transport.BaseResponse
	if err := c.transport.Delete(ctx, "/domain/"+domain+"/cancelation", &out); err != nil {
		return nil, err
	}
	return &out, nil
}
