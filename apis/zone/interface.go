package zone

import (
	"context"
)

// Client represents the functions implemented by this API
type Client interface {
	// All gets a list of all zones
	All(ctx context.Context) ([]*Zone, error)

	// Create creates a new zone
	Create(ctx context.Context, zone Zone) (*Response, error)

	// StreamUpdate updates the zone with given records
	StreamUpdate(ctx context.Context, zone Zone, adds []*ResourceRecord, removes []*ResourceRecord) (*Response, error)
}
