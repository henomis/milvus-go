package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorQuery struct {
	CollectionName string   `json:"collectionName"`
	Filter         string   `json:"filter"`
	OutputFields   []string `json:"outputFields,omitempty"`
	Limit          *string  `json:"limit,omitempty"`
	Offset         *string  `json:"offset,omitempty"`
}

func (c *VectorQuery) Path() (string, error) {
	return "/v1/vector/query", nil
}

func (c *VectorQuery) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *VectorQuery) ContentType() string {
	return ContentTypeApplicationJSON
}
