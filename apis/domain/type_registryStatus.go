package domain

// RegistryStatus represents the status at the registry
type RegistryStatus string

const (
	RegistryStatusActive     RegistryStatus = "ACTIVE"
	RegistryStatusHold       RegistryStatus = "HOLD"
	RegistryStatusLock       RegistryStatus = "LOCK"
	RegistryStatusHoldLock   RegistryStatus = "HOLD_LOCK"
	RegistryStatusAuto       RegistryStatus = "AUTO"
	RegistryStatusLockOwner  RegistryStatus = "LOCK_OWNER"
	RegistryStatusLockUpdate RegistryStatus = "LOCK_UPDATE"
	RegistryStatusPending    RegistryStatus = "PENDING"
	RegistryStatusNone       RegistryStatus = "NONE"
)
