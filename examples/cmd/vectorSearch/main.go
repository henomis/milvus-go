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

	resp := &response.VectorSearch{}
	err := client.VectorSearch(
		context.Background(),
		&request.VectorSearch{
			CollectionName: "test",
			Vector:         []float64{0.1, 0.2, 0.3, 0.4},
			OutputFields: []string{
				request.DefaultVectorField,
				request.DefaultPrimaryField,
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
