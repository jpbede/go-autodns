package zone

type ZoneStream struct {
	Adds    []*ResourceRecord `json:"adds"`
	Removes []*ResourceRecord `json:"rems"`
}
