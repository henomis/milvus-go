package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorDelete struct {
	CollectionName string   `json:"collectionName"`
	ID             []uint64 `json:"id"`
}

func (c *VectorDelete) Path() (string, error) {
	return "/v1/vector/delete", nil
}

func (c *VectorDelete) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *VectorDelete) ContentType() string {
	return ContentTypeApplicationJSON
}
