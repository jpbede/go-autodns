package contact

import (
	"context"
	"github.com/jpbede/go-autodns/apis/job"
	"github.com/jpbede/go-autodns/internal/transport"
)

// Client represents the functions implemented by this API
type Client interface {
	// All gets a list of all contacts
	All(ctx context.Context) ([]Contact, error)

	// Info gets information for a specific contact
	Info(contactID int, ctx context.Context) (*Contact, error)

	// Create creates a new contact
	Create(contact Contact, ctx context.Context) (*job.Response, error)

	// Update updates a given contact
	Update(contact Contact, ctx context.Context) (*job.Response, error)

	// UpdateComment updates a given contact comment
	UpdateComment(contact Contact, ctx context.Context) (*job.Response, error)

	// Delete deletes a contact
	Delete(contactID int, ctx context.Context) (*transport.BaseResponse, error)
}
