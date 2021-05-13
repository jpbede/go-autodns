package zone

import (
	"context"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) StreamUpdate(ctx context.Context, zone Zone, adds []*ResourceRecord, removes []*ResourceRecord) (*Response, error) {
	var out Response

	updates := ZoneStream{
		Adds:    adds,
		Removes: removes,
	}

	if err := c.transport.Post(ctx, "/zone/"+zone.Origin+"/_stream", &out, transport.WithJSONRequestBody(updates)); err != nil {
		return nil, err
	}
	return &out, nil
}
