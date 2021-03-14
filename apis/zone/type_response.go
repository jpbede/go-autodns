package zone

import "github.com/jpbede/go-autodns/internal/transport"

// Response represents the zone responses given by async tasks
type Response struct {
	transport.BaseResponse

	Data []*Zone
}
