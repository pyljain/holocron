package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "holocron",
	Short: "Holocron is the most popular and fastest vector store known to mankind",
	Long:  `Holocron is the most popular and fastest vector store known to mankind`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
