package response

import (
	"encoding/json"
	"io"
)

type CollectionDescribe struct {
	Response
	Data CollectionDescribeData `json:"data"`
}

type Fields struct {
	AutoID      bool   `json:"autoId"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PrimaryKey  bool   `json:"primaryKey"`
	Type        string `json:"type"`
}

type Indexes struct {
	FieldName  string `json:"fieldName"`
	IndexName  string `json:"indexName"`
	MetricType string `json:"metricType"`
}

type CollectionDescribeData struct {
	CollectionName     string    `json:"collectionName"`
	Description        string    `json:"description"`
	EnableDynamicField bool      `json:"enableDynamicField"`
	Fields             []Fields  `json:"fields"`
	Indexes            []Indexes `json:"indexes"`
	Load               string    `json:"load"`
	ShardsNum          int64     `json:"shardsNum"`
}

func (c *CollectionDescribe) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *CollectionDescribe) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
