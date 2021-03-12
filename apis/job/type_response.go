package job

import "github.com/jpbede/go-autodns/internal/transport"

// Response represents the job responses given by async tasks
type Response struct {
	transport.BaseResponse

	Data []*Job
}
