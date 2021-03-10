package domain

import (
	"context"
)

// Client represents the functions implemented by this API
type Client interface {
	All(ctx context.Context) ([]Domain, error)
	Info(domain string, ctx context.Context) (*Domain, error)
}
