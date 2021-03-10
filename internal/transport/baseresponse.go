package transport

type BaseResponse struct {
	StID string `json:"stid"`

	/*Messages []struct {
	  Text     string        `json:"text"`
	  Messages []interface{} `json:"messages"`
	  Objects  []struct {
	    Type  string `json:"type"`
	    Value string `json:"value"`
	  } `json:"objects"`
	  Code   string `json:"code"`
	  Status string `json:"status"`
	} `json:"messages"`*/
	Status struct {
		Code string `json:"code"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"status"`
	Object struct {
		Type    string `json:"type"`
		Value   string `json:"value"`
		Summary int    `json:"summary"`
		Data    struct {
		} `json:"data"`
	} `json:"object"`

	//CtID string `json:"ctid"`
}
