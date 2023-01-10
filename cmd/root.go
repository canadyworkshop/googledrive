package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"os"
)

var cliOps = struct {
	UserKeyFilePath string
}{
	UserKeyFilePath: "./.googledrive.key",
}

var driveService *drive.Service

var rootCmd = &cobra.Command{
	Use:              "googledrive",
	Short:            "Provides CLI access to Google Drive.",
	PersistentPreRun: setupDrive,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cliOps.UserKeyFilePath, "user-key-file", "u", cliOps.UserKeyFilePath, "The key file for the Google service account to use.")
}

func setupDrive(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	driveService, err = drive.NewService(
		ctx,
		option.WithCredentialsFile(cliOps.UserKeyFilePath))
	if err != nil {
		fmt.Printf("Failed to create client: %s\n", err)
		os.Exit(1)
	}

}
