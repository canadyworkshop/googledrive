package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var FileListCmdVars = struct {
	q string
	a bool
}{}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists files accessible to the connected account.",
	Run: func(cmd *cobra.Command, args []string) {
		driveService := mustConnectToDriveOrFail(RootVars.SACredsPath)

		// Building call based on args.

		fileList, err := driveService.Files.List().Do()
		if err != nil {
			log.Error(err.Error())
			fmt.Printf("failed to list to files: %s", err.Error())

			os.Exit(1)
		}
		for _, f := range fileList.Files {
			fmt.Println(f.Name)
		}
	},
}

func init() {
	fileCmd.AddCommand(fileListCmd)

	fileListCmd.Flags().StringVarP(&FileListCmdVars.q, "query", "q", "", "The query to use.")
	fileListCmd.Flags().StringVarP(&FileListCmdVars.q, "query", "q", "", "The query to use.")

	//
	//fileCreateCmd.Flags().StringVarP(&vCloudUsernameCLIOpt, "username", "u", "", "The username to use when connecting to vCloud director. Will override the env if set.")
	//fileCreateCmd.Flags().StringVarP(&vCloudPasswordClIOpt, "password", "p", "", "The password to use when connecting to vCloud director. Will override the env if set.")
	//fileCreateCmd.Flags().StringVarP(&vCloudEnvCLIOpt, "env", "e", "", "The name of the environment to look at. If empty the env variable will be used.")
}
