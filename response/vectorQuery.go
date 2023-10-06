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
	return json.NewDecoder(body).Decode(c)
}
