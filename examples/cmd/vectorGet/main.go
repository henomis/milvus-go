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

	resp := &response.VectorGet{}
	err := client.VectorGet(
		context.Background(),
		&request.VectorGet{
			CollectionName: "test",
			ID:             []uint64{444778967793664883},
			OutputFields: []string{
				request.DefaultVectorField,
				request.DefaultVectorField,
				"key",
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)
	for _, v := range resp.Data {
		fmt.Printf("id: %d\n", v.ID())
		fmt.Printf("vector: %v\n", v.Vector())
		fmt.Printf("key: %s\n", v["key"].(string))
	}

}
