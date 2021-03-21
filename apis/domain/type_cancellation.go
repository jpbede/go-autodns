package domain

type CancellationType string

const (
	CancellationTypeDelete  CancellationType = "DELETE"
	CancellationTypeTransit CancellationType = "TRANSIT"
	CancellationTypePreAck  CancellationType = "PREACK"
)

type CancellationExecution string

const (
	CancellationExecutionDate   CancellationExecution = "DATE"
	CancellationExecutionExpire CancellationExecution = "EXPIRE"
	CancellationExecutionNow    CancellationExecution = "NOW"
)

type Cancellation struct {
	Created          string                `json:"created"`
	Updated          string                `json:"updated"`
	Domain           string                `json:"domain"`
	RegistryWhen     string                `json:"registryWhen"`
	GainingRegistrar string                `json:"gainingRegistrar"`
	Disconnect       bool                  `json:"disconnect"`
	Notice           string                `json:"notice"`
	LogID            int                   `json:"logId"`
	RegistryStatus   RegistryStatus        `json:"registryStatus"`
	Type             CancellationType      `json:"type"`
	Execution        CancellationExecution `json:"execution"`
}
