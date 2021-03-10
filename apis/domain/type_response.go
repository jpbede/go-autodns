package domain

import "github.com/jpbede/go-autodns/internal/transport"

type DomainResponse struct {
	transport.BaseResponse

	Data []Domain `json:"data"`
}
