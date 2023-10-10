package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type CollectionCreate struct {
	DBName         *string `json:"dbName,omitempty"`
	CollectionName string  `json:"collectionName"`
	Dimension      int64   `json:"dimension"`
	MetricType     *Metric `json:"metricType"`
	PrimaryField   *string `json:"primaryField,omitempty"`
	VectorField    *string `json:"vectorField,omitempty"`
	Description    *string `json:"description,omitempty"`
}

type Metric string

const (
	MetricL2             Metric = "L2"
	MetricIp             Metric = "IP"
	MetricHamming        Metric = "HAMMING"
	MetricJaccard        Metric = "JACCARD"
	MetricTanimoto       Metric = "TANIMOTO"
	MetricSubstructure   Metric = "SUBSTRUCTURE"
	MetricSuperstructure Metric = "SUPERSTRUCTURE"
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
