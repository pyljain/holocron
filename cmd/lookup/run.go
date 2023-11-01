package lookup

import (
	"holocron/pkg/lookup"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the lookup service",
	Long:  "Starts the lookup service",
	RunE: func(cmd *cobra.Command, args []string) error {
		svr, err := lookup.NewServer(bucket)
		if err != nil {
			return err
		}

		err = svr.Start(port)
		if err != nil {
			return err
		}
		return nil
	},
}

var (
	bucket string
	port   int32
)

func init() {
	RunCmd.Flags().StringVar(&bucket, "bucket", "", "bucket to connect to")
	RunCmd.Flags().Int32Var(&port, "port", 8989, "Port")
	RunCmd.MarkFlagRequired("bucket")
}
