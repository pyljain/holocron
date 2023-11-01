package lookup

import (
	"context"
	"encoding/json"
	"fmt"
	"holocron/pkg/distance"
	"holocron/pkg/priorityqueue"
	"holocron/pkg/proto"
	"holocron/pkg/pubsub"
	"holocron/pkg/storage"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	bucket string
	store  storage.Storage
	proto.UnimplementedLookupServer
}

func NewServer(bucket string) (*Server, error) {

	store, err := storage.NewGCS()
	if err != nil {
		return nil, err
	}

	return &Server{
		bucket: bucket,
		store:  store,
	}, nil
}

func (s *Server) Start(port int32) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	svr := grpc.NewServer()
	proto.RegisterLookupServer(svr, s)
	err = svr.Serve(lis)
	return err
}

func (s *Server) Query(ctx context.Context, req *proto.LookupQueryRequest) (*proto.LookupQueryResponse, error) {
	queue := priorityqueue.NewQueue(int(req.TopK))

	// List objects in a bucket for a specific collection
	objectNames, err := s.store.List(ctx, s.bucket, req.Collection)
	if err != nil {
		return nil, err
	}

	for _, name := range objectNames {
		fileContent, err := s.store.Read(ctx, s.bucket, req.Collection, name)
		if err != nil {
			return nil, err
		}

		record := pubsub.Event{}

		err = json.Unmarshal(fileContent, &record)
		if err != nil {
			return nil, err
		}

		// Calculate vector distance
		// log.Printf("record.Embedding for %s is %+v", name, record.Embedding)
		d := distance.CalculateCosineSimilarity(req.Embedding, record.Embedding)
		queue.Push(d, &record)
	}

	// Read objects
	embeddings := []*proto.Embedding{}
	for _, element := range queue.Get() {
		embeddings = append(embeddings, &proto.Embedding{
			Embedding: element.MatchedVector.Embedding,
			Metadata:  element.MatchedVector.Metadata,
		})
	}

	return &proto.LookupQueryResponse{
		Embeddings: embeddings,
	}, nil
}
