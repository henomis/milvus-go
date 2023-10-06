package response

import (
	"encoding/json"
	"io"
)

type VectorInsert struct {
	Response
	Data VectorInsertData `json:"data"`
}

type VectorInsertData struct {
	InsertCount int64   `json:"insertCount"`
	InsertIds   []int64 `json:"insertIds"`
}

func (c *VectorInsert) AcceptContentType() string {
	return AcceptContentTypeApplicationJSON
}

func (c *VectorInsert) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(c)
}
