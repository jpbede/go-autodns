package job

import (
	"context"
)

// Client represents the functions implemented by this API
type Client interface {
	// All gets a list of all running/pending jobs
	All(ctx context.Context) ([]*ObjectJob, error)

	// Info gets information for a specific job
	Info(jobID int, ctx context.Context) (*ObjectJob, error)
}
