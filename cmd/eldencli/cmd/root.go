package cmd

import (
	"github.com/jon4hz/eldencli/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "devel", // TODO: change
	Use:     "eldencli",
	Short:   "Search, Tarnished",
	Long:    "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return root()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func root() error {
	return tui.Start()
}
