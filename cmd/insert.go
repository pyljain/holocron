package cmd

import (
	"context"
	"fmt"
	"holocron/pkg/client"
	"holocron/pkg/proto"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	addr       string
	file       string
	collection string
)

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringVar(&addr, "addr", "localhost:9999", "gRPC server address")
	insertCmd.Flags().StringVar(&collection, "collection", "", "Collection to insert into")
	insertCmd.Flags().StringVarP(&file, "file", "f", "", "File with embeddings")
	_ = insertCmd.MarkFlagRequired("file")
	_ = insertCmd.MarkFlagRequired("collection")

}

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a record into holocron",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Inserting record")
		c := client.New(addr)
		err := c.Connect()
		if err != nil {
			return err
		}

		fileContents, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		f := File{}
		err = yaml.Unmarshal(fileContents, &f)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		request := proto.EmbeddingWithMetadataRequest{
			Collection: collection,
			Embedding:  f.Embedding,
			Metadata:   f.Metadata,
		}
		status, err := c.Insert(ctx, &request)
		if err != nil {
			return err
		}

		log.Printf("Status is %s", status)

		return nil
	},
}

type File struct {
	Embedding []float64         `yaml:"embedding"`
	Metadata  map[string]string `yaml:"metadata"`
}
