package job

import (
	"go.bnck.me/autodns/internal/transport"
	"time"
)

type ObjectJob struct {
	Job            *Job                             `json:"job"`
	Object         *transport.ResponseObject        `json:"object"`
	NICComLogs     []*NICComLog                     `json:"niccomLogs"`
	Authentication []*CertAuthenticationStatus      `json:"authentication"`
	DCVAuth        []*DomainControlValidationStatus `json:"dcvAuth"`
}

type NICComLog struct {
	JobID            int       `json:"jobId"`
	VertexID         int       `json:"vertexId"`
	Name             string    `json:"name"`
	Task             string    `json:"task"`
	NicTransactionID string    `json:"nicTransactionId"`
	Source           string    `json:"source"`
	Text             string    `json:"text"`
	Created          time.Time `json:"created"`
}

type CertAuthenticationStatus struct {
	Status string `json:"status"`
	Step   string `json:"step"`
}

type DomainControlValidationStatus struct {
	Status string `json:"status"`
	Domain string `json:"domain"`
}
