package transport

type ResponseStatus struct {
	Code string `json:"code"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type ResponseObject struct {
	Type    string `json:"type"`
	Value   string `json:"value"`
	Summary int    `json:"summary"`
}

type MessageObject struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Message struct {
	Text     string          `json:"text"`
	Messages []string        `json:"messages"`
	Objects  []MessageObject `json:"objects"`
	Code     string          `json:"code"`
	Status   string          `json:"status"`
}

type BaseResponse struct {
	StID     string         `json:"stid"`
	Messages []Message      `json:"messages"`
	Status   ResponseStatus `json:"status"`
	Object   ResponseObject `json:"object"`
	CtID     string         `json:"ctid"`
}
