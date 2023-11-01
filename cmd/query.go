package cmd

import (
	"context"
	"holocron/pkg/client"
	"holocron/pkg/proto"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	topK int32
)

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVar(&addr, "addr", "localhost:9999", "gRPC server address")
	queryCmd.Flags().StringVar(&collection, "collection", "", "Collection to insert into")
	queryCmd.Flags().StringVarP(&file, "file", "f", "", "File with embeddings")
	queryCmd.Flags().Int32VarP(&topK, "topk", "k", 2, "Pass the number of results to return")
	_ = queryCmd.MarkFlagRequired("file")
	_ = queryCmd.MarkFlagRequired("collection")

}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run queries against Holocron to find aprroximate nearest neighbours",
	Long:  "Run queries against Holocron to find aprroximate nearest neighbours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("In query")

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

		resp, err := c.Query(ctx, &proto.QueryRequest{
			Embedding:  f.Embedding,
			Collection: collection,
			TopK:       topK,
		})

		if err != nil {
			return err
		}

		respEmbeddings := resp.Embeddings
		for _, e := range respEmbeddings {
			log.Printf("Response returned with embeddings %+v", e.Metadata)
		}

		return nil
	},
}
