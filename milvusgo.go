package milvusgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/henomis/milvus-go/request"
	"github.com/henomis/milvus-go/response"
	"github.com/henomis/restclientgo"
)

type Client struct {
	restClient *restclientgo.RestClient
	endpoint   string
}

func New(endpoint, username, password string) *Client {
	restClient := restclientgo.New(endpoint)

	token := fmt.Sprintf("%s:%s", username, password)

	if len(token) > 0 {
		restClient.SetRequestModifier(func(req *http.Request) *http.Request {
			req.Header.Set("Authorization", "Bearer "+token)
			return req
		})
	}
	return &Client{
		endpoint:   endpoint,
		restClient: restClient,
	}
}

func (c *Client) CollectionCreate(
	ctx context.Context,
	req *request.CollectionCreate,
	res *response.CollectionCreate,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) CollectionDescribe(
	ctx context.Context,
	req *request.CollectionDescribe,
	res *response.CollectionDescribe,
) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) CollectionDrop(
	ctx context.Context,
	req *request.CollectionDrop,
	res *response.CollectionDrop,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) CollectionList(
	ctx context.Context,
	req *request.CollectionList,
	res *response.CollectionList,
) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) VectorInsert(
	ctx context.Context,
	req *request.VectorInsert,
	res *response.VectorInsert,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) VectorGet(
	ctx context.Context,
	req *request.VectorGet,
	res *response.VectorGet,
) error {
	return c.restClient.Post(ctx, req, res)
}
