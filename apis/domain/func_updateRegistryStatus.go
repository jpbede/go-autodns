package domain

import (
	"context"
	"go.bnck.me/autodns/apis/job"
	"go.bnck.me/autodns/internal/transport"
)

func (c *client) UpdateRegistryStatus(domain Domain, regStatus RegistryStatus, ctx context.Context) (*job.Response, error) {
	domain.RegistryStatus = regStatus

	var out job.Response
	if err := c.transport.Put(ctx, "/domain/"+domain.Name+"/_statusUpdate", &out, transport.WithJSONRequestBody(domain)); err != nil {
		return nil, err
	}
	return &out, nil
}
