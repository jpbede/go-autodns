package domain

import (
	"context"
	"go.bnck.me/autodns/apis/contact"
	"go.bnck.me/autodns/apis/job"
	"go.bnck.me/autodns/internal/transport"
)

// Client represents the functions implemented by this API
type Client interface {
	// All gets a list of all registered domains
	All(ctx context.Context) ([]*Domain, error)

	// Info gets information for a specific domain
	Info(domain string, ctx context.Context) (*Domain, error)

	// Create orders a new domain
	Create(domain Domain, ctx context.Context) (*job.Response, error)

	// ChangeOwner changes the OwnerC of a given domain
	ChangeOwner(domain Domain, ownerC *contact.Contact, ctx context.Context) (*job.Response, error)

	// Update updates a given domain
	Update(domain Domain, ctx context.Context) (*job.Response, error)

	// UpdateComment updates a given domain
	UpdateComment(domain Domain, ctx context.Context) (*job.Response, error)

	// UpdateRegistryStatus updates the registry status a given domain
	UpdateRegistryStatus(domain Domain, regStatus RegistryStatus, ctx context.Context) (*job.Response, error)

	// Cancel cancels a domain
	Cancel(domain Domain, ctx context.Context) (*Cancellation, error)

	// DeleteCancellation gets all existing cancellations
	GetCancellations(ctx context.Context) ([]*Cancellation, error)

	// DeleteCancellation removes a existing cancellation for a domain
	DeleteCancellation(domain string, ctx context.Context) (*transport.BaseResponse, error)

	// Transfer transfers a domain
	Transfer(domain Domain, ctx context.Context) (*job.Response, error)

	// CreateAuthInfo1 creates an AuthInfo 1 for the specified domain.
	CreateAuthInfo1(domain Domain, ctx context.Context) (*Response, error)

	// SendAuthInfo1ToOwnerC sends the AuthInfo 1 for the specified domain to the OwnerC
	SendAuthInfo1ToOwnerC(domain Domain, ctx context.Context) (*transport.BaseResponse, error)

	// DeleteAuthInfo1 deletes an AuthInfo 1 for the specified domain.
	DeleteAuthInfo1(domain Domain, ctx context.Context) (*transport.BaseResponse, error)
}
