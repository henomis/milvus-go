package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type CollectionDrop struct {
	CollectionName string `json:"collectionName"`
}

func (c *CollectionDrop) Path() (string, error) {
	return "/v1/vector/collections/drop", nil
}

func (c *CollectionDrop) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *CollectionDrop) ContentType() string {
	return ContentTypeApplicationJSON
}
