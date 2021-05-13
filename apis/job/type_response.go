package job

import "go.bnck.me/autodns/internal/transport"

// Response represents the job responses given by async tasks
type Response struct {
	transport.BaseResponse

	Data []*Job
}
