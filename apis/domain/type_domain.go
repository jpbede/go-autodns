package domain

import (
	"go.bnck.me/autodns/apis/contact"
	"net"
)

// Domain represents a InternetX domain
type Domain struct {
	Created               string             `json:"created,omitempty"`
	Updated               string             `json:"updated,omitempty"`
	DomainCreated         string             `json:"domainCreated,omitempty"`
	Payable               string             `json:"payable,omitempty"`
	Expire                string             `json:"expire,omitempty"`
	RemoveCancellation    bool               `json:"removeCancelation,omitempty"`
	CancellationStatus    CancellationStatus `json:"cancelationStatus,omitempty"`
	AutoRenewStatus       AutoRenewStatus    `json:"autoRenewStatus,omitempty"`
	NICMemberLabel        string             `json:"nicMemberLabel,omitempty"`
	RegistryStatus        RegistryStatus     `json:"registryStatus,omitempty"`
	RegistrarStatusReason string             `json:"registrarStatusReason,omitempty"`
	AuthInfo              string             `json:"authinfo,omitempty"`
	Name                  string             `json:"name"`
	IDN                   string             `json:"idn,omitempty"`
	Nameserver            []*Nameserver      `json:"nameServers,omitempty"`
	DNSSEC                bool               `json:"dnssec,omitempty"`
	DNSSecData            []*DNSSecData      `json:"dnssecData,omitempty"`
	OwnerC                *contact.Contact   `json:"ownerc,omitempty"`
	AdminC                *contact.Contact   `json:"adminc,omitempty"`
	TechC                 *contact.Contact   `json:"techc,omitempty"`
	ZoneC                 *contact.Contact   `json:"zonec,omitempty"`
	Comment               string             `json:"comment,omitempty"`
}

// Nameserver represents a nameserver for a domain
type Nameserver struct {
	Name        string   `json:"name,omitempty"`
	TTL         int64    `json:"ttl,omitempty"`
	IPAddresses []net.IP `json:"ipAddresses,omitempty"`
}

// DNSSecData represents a DNSSec key material submitted to registry
type DNSSecData struct {
	Algorithm int    `json:"algorithm"`
	Flags     int    `json:"flags"`
	Protocol  int    `json:"protocol"`
	PublicKey string `json:"publicKey"`
}
