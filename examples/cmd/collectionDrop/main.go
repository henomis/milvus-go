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

	resp := &response.CollectionDrop{}
	err := client.CollectionDrop(
		context.Background(),
		&request.CollectionDrop{
			CollectionName: "test",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
