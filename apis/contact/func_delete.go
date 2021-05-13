package contact

import (
	"context"
	"go.bnck.me/autodns/internal/transport"
	"strconv"
)

func (c *client) Delete(contactID int, ctx context.Context) (*transport.BaseResponse, error) {
	var out transport.BaseResponse
	if err := c.transport.Delete(ctx, "/contact/"+strconv.Itoa(contactID), &out); err != nil {
		return nil, err
	}
	return &out, nil
}
