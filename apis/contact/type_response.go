package contact

import "github.com/jpbede/go-autodns/internal/transport"

// Response response of the /contact endpoints
type Response struct {
	transport.BaseResponse

	Data []Contact `json:"data"`
}
