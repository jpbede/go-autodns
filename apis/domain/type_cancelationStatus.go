package domain

// CancellationStatus status of a domain cancellation
type CancellationStatus string

const (
	CancellationStatusDelete        CancellationStatus = "DELETE"
	CancellationStatusDeleteExpire  CancellationStatus = "DELETE_EXPIRE"
	CancellationStatusTransit       CancellationStatus = "TRANSIT"
	CancellationStatusTransitExpire CancellationStatus = "TRANSIT_EXPIRE"
	CancellationStatusPreAck        CancellationStatus = "PREACK"
	CancellationStatusPreAckExpire  CancellationStatus = "PREACK_EXPIRE"
)
