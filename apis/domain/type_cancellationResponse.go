package domain

import "go.bnck.me/autodns/internal/transport"

// Response response of the /domain endpoints
type CancellationResponse struct {
	transport.BaseResponse

	Data []*Cancellation `json:"data"`
}
