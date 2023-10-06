package response

import (
	"encoding/json"
	"io"
)

type CollectionDrop struct {
	Response
	Data any `json:"data"`
}

func (c *CollectionDrop) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionDrop) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
