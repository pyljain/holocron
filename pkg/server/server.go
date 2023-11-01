package server

import (
	"context"
	"fmt"
	"holocron/pkg/proto"
	"holocron/pkg/pubsub"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedHolocronServer
	publisher    pubsub.Publisher
	lookupClient proto.LookupClient
}

func New(publisher pubsub.Publisher, lookupClient proto.LookupClient) *Server {
	return &Server{
		publisher:    publisher,
		lookupClient: lookupClient,
	}
}

func (s *Server) Start(port int32) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	svr := grpc.NewServer()
	proto.RegisterHolocronServer(svr, s)
	err = svr.Serve(lis)
	return err
}

func (s *Server) Insert(ctx context.Context, req *proto.EmbeddingWithMetadataRequest) (*proto.InsertStatus, error) {
	err := s.publisher.Publish(req.Collection, req.Embedding, req.Metadata)
	if err != nil {
		return nil, err
	}
	return &proto.InsertStatus{
		Message: "Success",
	}, nil
}

func (s *Server) Query(ctx context.Context, req *proto.QueryRequest) (*proto.QueryResponse, error) {
	log.Printf("Query in proxy")
	queryResp, err := s.lookupClient.Query(ctx, &proto.LookupQueryRequest{
		Collection: req.Collection,
		Embedding:  req.Embedding,
		TopK:       req.TopK,
	})

	if err != nil {
		return nil, err
	}

	qr := proto.QueryResponse{}

	for _, re := range queryResp.Embeddings {
		qr.Embeddings = append(qr.Embeddings, &proto.EmbeddingWithMetadataRequest{
			Collection: req.Collection,
			Embedding:  re.Embedding,
			Metadata:   re.Metadata,
		})
	}

	return &qr, nil
}
