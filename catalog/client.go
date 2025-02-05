package catalog

import (
	"context"

	"github.com/mohammedyunus2002/go-grpc-graphql-microservices/catalog/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.CatalogServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewCataogServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	c.service.PostProduct
}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error) {

}

func (c *Client) GetProducts(ctx context.Context, id string) (*Product, error) {

}
