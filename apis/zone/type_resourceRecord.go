package zone

type ResourceRecord struct {
	Name  string `json:"name,omitempty"`
	TTL   int    `json:"ttl,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Pref  int    `json:"pref,omitempty"`
	Raw   string `json:"raw,omitempty"`
}
