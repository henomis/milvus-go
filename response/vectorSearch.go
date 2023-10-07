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
	dec := json.NewDecoder(body)
	dec.UseNumber()
	return dec.Decode(c)
}
