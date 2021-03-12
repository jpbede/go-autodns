package domain

import "github.com/jpbede/go-autodns/internal/transport"

// DomainResponse response of the /domain endpoints
type DomainResponse struct {
	transport.BaseResponse

	Data []Domain `json:"data"`
}
