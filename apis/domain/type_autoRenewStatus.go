package domain

// AutoRenewStatus represents the status of the auto renew
type AutoRenewStatus string

const (
	AutoRenewStatusOn   AutoRenewStatus = "TRUE"
	AutoRenewStatusOff  AutoRenewStatus = "FALSE"
	AutoRenewStatusOnce AutoRenewStatus = "ONCE"
)
