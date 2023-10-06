# Unofficial Milvus Go SDK


[![GoDoc](https://godoc.org/github.com/henomis/milvus-go?status.svg)](https://godoc.org/github.com/henomis/milvus-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/milvus-go)](https://goreportcard.com/report/github.com/henomis/milvus-go) [![GitHub release](https://img.shields.io/github/release/henomis/milvus-go.svg)](https://github.com/henomis/milvus-go/releases)

This is [Milvus](https://milvus.tech/)'s **unofficial** Go client, designed to enable you to use Milvus's services easily from your own applications.

## Milvus

[Milvus](https://milvus.tech/) is a vector database that allows you to build high-performance vector search applications.

## API support

### collections
- ✅ list
- ✅ create
- ✅ collect info
- ✅ update
- ✅ delete
- ❌ update aliases
- ✅ create index
- ❌ delete index
- ❌ cluster info
- ❌ update cluster setup
- ❌ list aliases
- ❌ recover from uploaded snapshot
- ❌ recover from snapshot
- ❌ create snapshot
- ❌ delete snapshot
- ❌ download snapshot


### points 
- ✅ get point
- ✅ get points
- ✅ upsert points
- ✅ delete points
- ❌ update vectors
- ❌ delete vectors
- ❌ set payload
- ❌ overwrite payload
- ❌ delete payload
- ❌ clear payload
- ❌ scroll payload
- ✅ search points
- ❌ search batch points
- ❌ search point groups
- ❌ recommend points
- ❌ recommend batch points
- ❌ recommend point groups
- ❌ count points

### cluster
- ❌ cluster status info
- ❌ tries to recover current peer Raft state
- ❌ remove peer
- ❌ collection cluster info
- ❌ update collection cluster setup

### snapshots
- ❌ recover from uploaded snapshot
- ❌ recover from snapshot
- ❌ list collection snapshots
- ❌ create collection snapshot
- ❌ delete collection snapshot
- ❌ download collection snapshot
- ❌ list storage snapshots
- ❌ create storage snapshot
- ❌ delete storage snapshot
- ❌ download storage snapshot

### service
- ❌ collect telemetry data
- ❌ collect Prometheus metrics data
- ❌ set lock options
- ❌ get lock options


## Getting started

### Installation

You can load milvus-go into your project by using:
```
go get github.com/henomis/milvus-go
```

### Run Milvus

You can run Milvus using Docker:
```shell
docker run -p 6333:6333 --name milvus --rm -v $(pwd)/config.yaml:/milvus/config/production.yaml milvus/milvus
```

config.yaml file:
```yaml
service:
  api_key: secret-api-key
```

Please refer to the [official documentation](https://milvus.tech/) for more information about Milvus.

### Configuration

The only thing you need to start using Milvus's APIs is the API key. Copy and paste it in the corresponding place in the code, select the API and the parameters you want to use, and that's it.


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

	client := milvusgo.New("http://localhost:6333", "secret-api-key")

	onDisk := true
	resp := &response.CollectionCreate{}
	err := client.CollectionCreate(
		context.Background(),
		&request.CollectionCreate{
			CollectionName: "test",
			Vectors: request.VectorsParams{
				Size:     4,
				Distance: request.DistanceCosine,
				OnDisk:   &onDisk,
			},
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