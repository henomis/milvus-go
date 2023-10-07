package response

import (
	"encoding/json"
	"io"
)

type VectorQuery struct {
	Response
	Data []VectorData `json:"data"`
}

func (c *VectorQuery) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *VectorQuery) Decode(body io.Reader) error {
	dec := json.NewDecoder(body)
	dec.UseNumber()
	return dec.Decode(c)
}
