package contact

import "go.bnck.me/autodns/internal/transport"

// Response response of the /contact endpoints
type Response struct {
	transport.BaseResponse

	Data []Contact `json:"data"`
}
