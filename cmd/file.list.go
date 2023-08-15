package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var FileListCmdVars = struct {
	query          string
	allDrives      bool
	driveID        string
	fields         string
	includeLabels  string
	orderBy        string
	pageSize       int64
	paged          bool
	pageToken      string
	listLongFormat bool
}{}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists files accessible to the connected account.",
	Run: func(cmd *cobra.Command, args []string) {
		driveService := mustConnectToDriveOrFail(RootVars.SACredsPath)

		// Building call based on args.
		call := driveService.Files.List().SupportsAllDrives(true)

		call = call.IncludeItemsFromAllDrives(FileListCmdVars.allDrives)

		if FileListCmdVars.query != "" {
			call = call.Q(FileListCmdVars.query)
		}

		if FileListCmdVars.driveID != "" {
			call = call.DriveId(FileListCmdVars.driveID)
		}

		if FileListCmdVars.fields != "" {
			call = call.DriveId(FileListCmdVars.fields)
		}

		if FileListCmdVars.includeLabels != "" {
			call = call.DriveId(FileListCmdVars.includeLabels)
		}

		if FileListCmdVars.orderBy != "" {
			call = call.OrderBy(FileListCmdVars.orderBy)
		}

		if FileListCmdVars.pageSize > 0 {
			call = call.PageSize(FileListCmdVars.pageSize)
		}

		if FileListCmdVars.paged && FileListCmdVars.pageToken != "" {
			call = call.PageToken(FileListCmdVars.pageToken)
		}

		fileList, err := call.Do()
		if err != nil {
			log.Error(err.Error())
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if FileListCmdVars.listLongFormat {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			for _, f := range fileList.Files {
				fmt.Fprintf(w, "%s\t%s\t%d\t%s\t\n", f.Id, f.Name, f.Size, f.MimeType)

			}
			w.Flush()
		} else {
			for _, f := range fileList.Files {
				fmt.Println(f.Name)
			}
		}

	},
}

func init() {
	fileCmd.AddCommand(fileListCmd)

	fileListCmd.Flags().StringVar(&FileListCmdVars.query, "query", "", "The query to use.")
	fileListCmd.Flags().BoolVar(&FileListCmdVars.allDrives, "all-drives", true, "Denotes that all drives should be searched.")
	fileListCmd.Flags().StringVar(&FileListCmdVars.driveID, "drive-id", "", "The ID of the drive to search.")
	fileListCmd.Flags().StringVar(&FileListCmdVars.fields, "fields", "", "An optional list of fields to limit the returned data to.")
	fileListCmd.Flags().StringVar(&FileListCmdVars.includeLabels, "include-labels", "", "A comma separated list of IDs for labels that should be returned if on the file.")
	fileListCmd.Flags().StringVar(&FileListCmdVars.orderBy, "order-by", "", "A comma separated list of fields to order the results by.")
	fileListCmd.Flags().Int64Var(&FileListCmdVars.pageSize, "page-size", 0, "The maximum number of files to return per API request.")
	fileListCmd.Flags().BoolVar(&FileListCmdVars.paged, "paged", false, "Enables paged results.")
	fileListCmd.Flags().StringVar(&FileListCmdVars.pageToken, "page-token", "", "The token for the next page when using the paged option.")
	fileListCmd.Flags().BoolVarP(&FileListCmdVars.listLongFormat, "list", "l", false, "List files in long format.")
}
