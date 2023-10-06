package response

import (
	"encoding/json"
	"io"
)

type VectorDelete struct {
	Response
	Data any `json:"data"`
}

func (c *VectorDelete) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *VectorDelete) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
