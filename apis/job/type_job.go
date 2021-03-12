package job

type Status string

const (
	StatusRunning  Status = "RUNNING"
	StatusSuccess  Status = "SUCCESS"
	StatusFailed   Status = "FAILED"
	StatusCanceled Status = "CANCELED"
	StatusSupport  Status = "SUPPORT"
	StatusDeferred Status = "DEFERRED"
	StatusNotSet   Status = "NOT_SET"
	StatusWait     Status = "WAIT"
)

type Job struct {
	ID        int     `json:"id"`
	Created   string  `json:"created"`
	Updated   string  `json:"updated"`
	Status    Status  `json:"status"`
	Events    []Event `json:"events"`
	SubStatus string  `json:"subStatus"`
	Execution string  `json:"execution"`
	Action    string  `json:"action"`
	SubType   string  `json:"subType"`
}

type Event struct {
	Vertex int32  `json:"vertex"`
	Type   string `json:"type"`
}
