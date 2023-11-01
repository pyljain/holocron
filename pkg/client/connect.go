package client

import (
	"holocron/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	serverAddress string
	proto.HolocronClient
}

func New(serverAddress string) *Client {
	return &Client{
		serverAddress: serverAddress,
	}
}

func (c *Client) Connect() error {
	conn, err := grpc.Dial(c.serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.HolocronClient = proto.NewHolocronClient(conn)
	return nil
}
