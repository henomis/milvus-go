package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type CollectionCreate struct {
	DBName         *string     `json:"dbName,omitempty"`
	CollectionName string      `json:"collectionName"`
	Dimension      int64       `json:"dimension"`
	MetricType     *MetricType `json:"metricType"`
	PrimaryField   *string     `json:"primaryField,omitempty"`
	VectorField    *string     `json:"vectorField,omitempty"`
	Description    *string     `json:"description,omitempty"`
}

type MetricType string

const (
	L2             MetricType = "L2"
	IP             MetricType = "IP"
	HAMMING        MetricType = "HAMMING"
	JACCARD        MetricType = "JACCARD"
	TANIMOTO       MetricType = "TANIMOTO"
	SUBSTRUCTURE   MetricType = "SUBSTRUCTURE"
	SUPERSTRUCTURE MetricType = "SUPERSTRUCTURE"
)

func (c *CollectionCreate) Path() (string, error) {
	return "/v1/vector/collections/create", nil
}

func (c *CollectionCreate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (c *CollectionCreate) ContentType() string {
	return ContentTypeApplicationJSON
}
