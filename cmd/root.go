package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmdOptions struct {
	Verbose bool
}

var rootCmd = &cobra.Command{
	Use:  "sync",
	Long: "Sync data from sources to targets",
}

// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	dieOnError("", err)
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.Verbose, "verbose", false, "Set to get more detailed output")
}
