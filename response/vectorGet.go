package response

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	DefaultVectorField  = "vector"
	DefaultPrimaryField = "id"
)

type VectorData map[string]any

func (c VectorData) ID() int64 {
	if id, isInt64 := c[DefaultPrimaryField].(int64); isInt64 {
		return id
	} else if id, isJSONNumber64 := c[DefaultPrimaryField].(json.Number); isJSONNumber64 {
		idInt64, _ := strconv.ParseInt(string(id), 10, 64)
		return idInt64
	}
	return 0
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
		} else if floatValue, isJSONNumber := v.(json.Number); isJSONNumber {
			floatSlice[i], _ = strconv.ParseFloat(string(floatValue), 64)
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
	dec := json.NewDecoder(body)
	dec.UseNumber()
	return dec.Decode(c)
}
