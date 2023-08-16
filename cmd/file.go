package cmd

import "github.com/spf13/cobra"

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Commands that act on files.",
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
