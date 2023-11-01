package cmd

import (
	"holocron/cmd/lookup"

	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Starts the lookup service",
	Long:  "Starts the lookup service",
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.AddCommand(lookup.RunCmd)
}
