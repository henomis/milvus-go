package main

import (
	"context"
	"fmt"

	milvusgo "github.com/henomis/milvus-go"
	"github.com/henomis/milvus-go/request"
	"github.com/henomis/milvus-go/response"
)

func main() {

	client := milvusgo.New("http://localhost:19530", "root", "Milvus")

	resp := &response.VectorInsert{}
	err := client.VectorInsert(
		context.Background(),
		&request.VectorInsert{
			CollectionName: "test",
			Data: []request.VectorData{
				{
					"key":                      "value",
					request.DefaultVectorField: []float32{0.1, 0.2, 0.3, 0.4},
				},
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
