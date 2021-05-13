package zone

import "go.bnck.me/autodns/internal/transport"

// Response represents the zone responses given by async tasks
type Response struct {
	transport.BaseResponse

	Data []*Zone
}
