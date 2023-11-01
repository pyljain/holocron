package ingestor

import (
	"context"
	"encoding/json"
	"fmt"
	"holocron/pkg/pubsub"
	"holocron/pkg/storage"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	natsURL string
	bucket  string
)

func init() {
	RunCmd.Flags().StringVar(&natsURL, "natsURL", "nats://localhost:4222", "The Nats URL")
	RunCmd.Flags().StringVar(&bucket, "bucket", "", "The bucket in which to store the data")
	RunCmd.MarkFlagRequired("bucket")
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run ingestor to index and persist embeddings",
	Long:  "Run ingestor to index and persist embeddings",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()

		natsConn, err := pubsub.NewNATS(natsURL)
		if err != nil {
			return err
		}
		defer natsConn.Drain()

		var store storage.Storage

		store, err = storage.NewGCS()
		if err != nil {
			return err
		}

		eventCh := make(chan pubsub.Event)
		defer close(eventCh)

		go func() {
			err = natsConn.Pull(eventCh)
			if err != nil {
				log.Printf("Error in pull %s", err)
				close(eventCh)
			}
		}()

		for e := range eventCh {
			eventBytes, err := json.Marshal(e)
			if err != nil {
				log.Printf("Error in marshal %s", err)
				continue
			}

			filename := fmt.Sprintf("%s/%s", e.Collection, uuid.NewString())

			err = store.Write(ctx, bucket, filename, eventBytes)
			if err != nil {
				log.Printf("Error in writing event to storage %s", err)
				continue
			}
		}

		return nil
	},
}
