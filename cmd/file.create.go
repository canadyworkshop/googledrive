package cmd

import (
	"github.com/spf13/cobra"
)

var fileCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a file by uploading it.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	fileCmd.AddCommand(fileCreateCmd)
	//
	//fileCreateCmd.Flags().StringVarP(&vCloudUsernameCLIOpt, "username", "u", "", "The username to use when connecting to vCloud director. Will override the env if set.")
	//fileCreateCmd.Flags().StringVarP(&vCloudPasswordClIOpt, "password", "p", "", "The password to use when connecting to vCloud director. Will override the env if set.")
	//fileCreateCmd.Flags().StringVarP(&vCloudEnvCLIOpt, "env", "e", "", "The name of the environment to look at. If empty the env variable will be used.")
}
