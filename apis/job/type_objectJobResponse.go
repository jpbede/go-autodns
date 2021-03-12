package job

import "github.com/jpbede/go-autodns/internal/transport"

// ObjectJobResponse represents a ObjectJob responses given by async tasks
type ObjectJobResponse struct {
	transport.BaseResponse

	Data []*ObjectJob `json:"data"`
}
