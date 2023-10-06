# Unofficial Milvus Go SDK


[![GoDoc](https://godoc.org/github.com/henomis/milvus-go?status.svg)](https://godoc.org/github.com/henomis/milvus-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/milvus-go)](https://goreportcard.com/report/github.com/henomis/milvus-go) [![GitHub release](https://img.shields.io/github/release/henomis/milvus-go.svg)](https://github.com/henomis/milvus-go/releases)

This is [Milvus](https://milvus.tech/)'s **unofficial** Go client, designed to enable you to use Milvus's services easily from your own applications.

## Milvus

[Milvus](https://milvus.tech/) is a vector database that allows you to build high-performance vector search applications.

## API support

### collections
- ✅ list
- ✅ create
- ✅ describe
- ✅ update
- ✅ drop

### vectors 
- ✅ get
- ✅ insert
- ✅ delete
- ✅ query
- ✅ search


## Getting started

### Installation

You can load milvus-go into your project by using:
```
go get github.com/henomis/milvus-go
```

### Run Milvus

You can run Milvus using Docker:
```shell
cd docker && docker-compose up -d
```


Please refer to the [official documentation](https://milvus.io/) for more information about Milvus.

### Configuration

The only thing you need to start using Milvus's is username and password. Copy and paste it in the corresponding place in the code, select the API and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
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
			Vector:         []float64{2, 3, 3, 4},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
```

## Who uses milvus-go?

* [LinGoose](https://github.com/henomis/lingoose) Go framework for building awesome LLM apps