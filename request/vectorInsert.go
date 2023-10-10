package request

import (
	"bytes"
	"encoding/json"
	"io"
)

const (
	DefaultVectorField   = "vector"
	DefaultPrimaryField  = "id"
	DefaultDistanceField = "distance"
)

type VectorInsert struct {
	CollectionName string       `json:"collectionName"`
	Data           []VectorData `json:"data"`
}

type VectorData map[string]any

func (c *VectorInsert) Path() (string, error) {
	return "/v1/vector/insert", nil
}

func (c *VectorInsert) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *VectorInsert) ContentType() string {
	return ContentTypeApplicationJSON
}
