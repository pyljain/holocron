package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GCS struct {
	client *storage.Client
}

func NewGCS() (Storage, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &GCS{client}, nil
}

func (g *GCS) Write(ctx context.Context, bucket string, filename string, data []byte) error {
	obj := g.client.Bucket(bucket).Object(filename)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, bytes.NewBuffer(data)); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (g *GCS) List(ctx context.Context, bucket string, collection string) ([]string, error) {
	bkt := g.client.Bucket(bucket)
	log.Printf("Searching GCS with prefix %s", collection)
	objIterator := bkt.Objects(ctx, &storage.Query{
		Prefix: collection,
	})

	objects := []string{}

	for {
		attrs, err := objIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %w", bucket, err)
		}

		objects = append(objects, attrs.Name)
	}

	log.Printf("Objects from GCS are %+v", objects)
	return objects, nil
}

func (g *GCS) Read(ctx context.Context, bucket string, collection string, filename string) ([]byte, error) {
	obj := g.client.Bucket(bucket).Object(filename)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Bucket(%q).Objects: %w", bucket, err)
	}

	buff := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buff, reader)
	if err != nil {
		return nil, fmt.Errorf("Bucket(%q).Objects: %w", bucket, err)
	}

	return buff.Bytes(), nil
}
