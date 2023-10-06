package response

import (
	"encoding/json"
	"io"
)

type CollectionList struct {
	Response
	Data []string `json:"data"`
}

func (c *CollectionList) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *CollectionList) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
