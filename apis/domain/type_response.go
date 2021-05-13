package domain

import "go.bnck.me/autodns/internal/transport"

// Response response of the /domain endpoints
type Response struct {
	transport.BaseResponse

	Data []*Domain `json:"data"`
}
