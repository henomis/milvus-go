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
	return AcceptContentTypeApplicationJSON
}

func (c *CollectionDrop) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
