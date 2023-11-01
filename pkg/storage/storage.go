package storage

import "context"

type Storage interface {
	Write(ctx context.Context, bucket string, filename string, data []byte) error
	List(ctx context.Context, bucket string, collection string) ([]string, error)
	Read(ctx context.Context, bucket string, collection string, filename string) ([]byte, error)
}
