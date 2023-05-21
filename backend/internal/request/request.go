package request

type Request struct {
	Entity   string         `json:"entity"`
	Data     map[string]any `json:"data"`
	Metadata Metadata       `json:"metadata"`
}

type GetRequest struct {
	Entity   string   `json:"entity"`
	Data     Data     `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	ID string `json:"id"`
}

type Metadata struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Response struct {
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
	Data   any      `json:"data,omitempty"`
	Total  int      `json:"total,omitempty"`
}
