package response

import (
	"encoding/json"
	"io"
)

type VectorSearch struct {
	Response
	Data []VectorData `json:"data"`
}

func (c *VectorSearch) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *VectorSearch) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
