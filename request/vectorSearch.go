package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorSearch struct {
	CollectionName string    `json:"collectionName"`
	Filter         *string   `json:"filter"`
	OutputFields   []string  `json:"outputFields,omitempty"`
	Limit          *string   `json:"limit"`
	Offset         *string   `json:"offset"`
	Vector         []float64 `json:"vector"`
}

func (c *VectorSearch) Path() (string, error) {
	return "/v1/vector/search", nil
}

func (c *VectorSearch) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *VectorSearch) ContentType() string {
	return ContentTypeApplicationJSON
}
