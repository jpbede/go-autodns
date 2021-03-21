package domain

import "github.com/jpbede/go-autodns/internal/transport"

// Response response of the /domain endpoints
type CancellationResponse struct {
	transport.BaseResponse

	Data []*Cancellation `json:"data"`
}
