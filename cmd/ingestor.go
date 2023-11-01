package cmd

import (
	"holocron/cmd/ingestor"

	"github.com/spf13/cobra"
)

var ingestorCmd = &cobra.Command{
	Use:   "ingestor",
	Short: "Starts the process of persisting vectors in object strorage",
	Long:  "Starts the process of persisting vectors in object strorage",
}

func init() {
	rootCmd.AddCommand(ingestorCmd)
	ingestorCmd.AddCommand(ingestor.RunCmd)
}
