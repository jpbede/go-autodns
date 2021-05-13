package zone

import (
	"go.bnck.me/autodns/apis/domain"
	"time"
)

type NameserverAction string

const (
	NameserverActionPrimary   NameserverAction = "PRIMARY"
	NameserverActionSecondary NameserverAction = "SECONDARY"
	NameserverActionComplete  NameserverAction = "COMPLETE"
	NameserverActionHidden    NameserverAction = "HIDDEN"
	NameserverActionNone      NameserverAction = "NONE"
	NameserverActionAuto      NameserverAction = "AUTO"
)

type Zone struct {
	Created               time.Time
	Updated               time.Time
	Origin                string
	IDN                   string
	DNSSec                bool
	NameserverGroup       string
	AllowTransfer         bool
	LogID                 int
	Comment               string
	DomainSafe            bool
	Source                string
	SourceVirtualHostname string
	Nameservers           []*domain.Nameserver
	WWWInclude            bool
	VirtualNameserver     bool
	FreeText              []string
	RoID                  int
	Grants                []string

	SOA             *SOA
	Main            *MainIP
	Action          NameserverAction
	ResourceRecords []*ResourceRecord
}

type MainIP struct {
	Address string
	TTL     int
}

type SOA struct {
	Refresh int
	Retry   int
	Expire  int
	TTL     int
	EMail   string
}
