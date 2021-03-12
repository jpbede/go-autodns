package contact

import (
	"context"
	"errors"
	"strconv"
)

func (c *client) Info(contactID int, ctx context.Context) (*Contact, error) {
	var out Response
	contactIDStr := strconv.Itoa(contactID)
	if err := c.transport.Get(ctx, "/contact/"+contactIDStr, &out); err != nil {
		return nil, err
	}
	if len(out.Data) > 0 {
		return &out.Data[0], nil
	} else {
		return nil, errors.New("no domain found")
	}
}
