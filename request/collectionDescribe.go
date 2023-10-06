package request

import (
	"fmt"
	"io"
)

type CollectionDescribe struct {
	CollectionName string `json:"-"`
}

func (c *CollectionDescribe) Path() (string, error) {
	return fmt.Sprintf("/v1/vector/collections/describe?collectionName=%s", c.CollectionName), nil
}

func (c *CollectionDescribe) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *CollectionDescribe) ContentType() string {
	return ""
}
