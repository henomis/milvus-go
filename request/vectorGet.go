package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorGet struct {
	CollectionName string   `json:"collectionName"`
	ID             []uint64 `json:"id"`
	OutputFields   []string `json:"outputFields,omitempty"`
}

func (c *VectorGet) Path() (string, error) {
	return "/v1/vector/get", nil
}

func (c *VectorGet) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *VectorGet) ContentType() string {
	return ContentTypeApplicationJSON
}
