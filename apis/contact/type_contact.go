package contact

// ContactType type of the contact
type ContactType string

const (
	ContactTypeRole   ContactType = "ROLE"
	ContactTypePerson ContactType = "PERSON"
	ContactTypeOrg    ContactType = "ORG"
)

// Contact represents a contact used for OwnerC, AdminC, TechC or ZoneC
type Contact struct {
	ID           int         `json:"id,omitempty"`
	Updated      string      `json:"updated,omitempty"`
	Created      string      `json:"created,omitempty"`
	Type         ContactType `json:"type,omitempty"`
	Alias        string      `json:"alias,omitempty"`
	Organization string      `json:"organization,omitempty"`
	Title        string      `json:"title,omitempty"`
	City         string      `json:"city,omitempty"`
	Country      string      `json:"country,omitempty"`
	State        string      `json:"state,omitempty"`
	EMail        string      `json:"email,omitempty"`
	SIP          string      `json:"sip,omitempty"`
	Remarks      []string    `json:"remarks,omitempty"`
	Firstname    string      `json:"fname,omitempty"`
	Lastname     string      `json:"lname,omitempty"`
	Address      []string    `json:"address,omitempty"`
	PostCode     string      `json:"pcode,omitempty"`
	Phone        string      `json:"phone,omitempty"`
	Fax          string      `json:"fax,omitempty"`
	Comment      string      `json:"comment,omitempty"`

	confirmOwnerConsent bool `json:"confirmOwnerConsent,omitempty"`
	domainsafe          bool `json:"domainsafe,omitempty"`
}
