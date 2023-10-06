package response

import (
	"encoding/json"
	"io"
)

type CollectionCreate struct {
	Response
	Data any `json:"data"`
}

func (c *CollectionCreate) AcceptContentType() string {
	return "application/json"
}

func (c *CollectionCreate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
