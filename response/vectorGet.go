package response

import (
	"encoding/json"
	"io"
)

const (
	DefaultVectorField  = "vector"
	DefaultPrimaryField = "id"
)

type VectorData map[string]any

func (c VectorData) ID() int64 {
	return int64(c[DefaultPrimaryField].(float64))
}

func (c VectorData) Vector() []float64 {
	interfaceSlice, isIterfaceSliceOk := c[DefaultVectorField].([]interface{})
	if !isIterfaceSliceOk {
		return []float64{}
	}

	floatSlice := make([]float64, len(interfaceSlice))
	for i, v := range interfaceSlice {
		if floatValue, isFloatValue := v.(float64); isFloatValue {
			floatSlice[i] = floatValue
		}
	}

	return floatSlice
}

type VectorGet struct {
	Response
	Data []VectorData `json:"data"`
}

func (c *VectorGet) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *VectorGet) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
