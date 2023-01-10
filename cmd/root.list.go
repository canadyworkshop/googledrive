package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists files in the accessible google drive.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test")

		files, err := driveService.Files.List().Do()
		if err != nil {
			fmt.Printf("Failed to retrieve files: %s\n", err)
			os.Exit(1)
		}

		for _, f := range files.Files {
			fmt.Printf("%s\n", f.Name)
			fmt.Println("  %d parents\n", len(f.Parents))
			for _, p := range f.Parents {
				fmt.Printf("   %s\n", p)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(rootListCmd)

}
