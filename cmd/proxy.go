package cmd

import (
	"fmt"
	"holocron/pkg/lookup"
	"holocron/pkg/pubsub"
	"holocron/pkg/server"

	"github.com/spf13/cobra"
)

var (
	natsURL   string
	lookupURL string
)

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().StringVar(&natsURL, "natsURL", "nats://localhost:4222", "NATS server URL")
	proxyCmd.Flags().StringVar(&lookupURL, "lookupURL", "localhost:8989", "NATS server URL")
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run the main holocron proxy",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting proxy")
		natsConn, err := pubsub.NewNATS(natsURL)
		if err != nil {
			return err
		}
		defer natsConn.Drain()

		lookupClient := lookup.NewClient(lookupURL)
		err = lookupClient.Connect()
		if err != nil {
			return err
		}

		s := server.New(natsConn, lookupClient)
		err = s.Start(9999)
		if err != nil {
			return err
		}

		return nil
	},
}
