package lookup

import (
	"holocron/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	proto.LookupClient
	addr string
}

func NewClient(addr string) *client {
	return &client{
		addr: addr,
	}
}

func (c *client) Connect() error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.LookupClient = proto.NewLookupClient(conn)
	return nil
}
