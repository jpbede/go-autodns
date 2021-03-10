package domain

import "net"

type Domain struct {
	AuthInfo   string        `json:"authinfo"`
	Name       string        `json:"name"`
	Nameserver []*Nameserver `json:"nameServers"`
}

type Nameserver struct {
	Name        string   `json:"name"`
	TTL         int64    `json:"ttl"`
	IPAddresses []net.IP `json:"ipAddresses"`
}
