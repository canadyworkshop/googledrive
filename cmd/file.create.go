package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"os"
	"strings"
)

var FileCreateCmdVars = struct {
	description      string
	includedLabelIDs string
	mimeType         string
	name             string
	parents          string
	filePath         string
}{}

var fileCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new file by uploading it.",
	Run: func(cmd *cobra.Command, args []string) {
		// Find and load file.
		file, err := os.Open(FileCreateCmdVars.filePath)
		if err != nil {
			fmt.Printf("Failed to find file %s\n", FileCreateCmdVars.filePath)
			log.Error(err.Error())
			os.Exit(1)
		}
		info, err := file.Stat()
		if err != nil {
			fmt.Printf("Failed to stat file %s\n", FileCreateCmdVars.filePath)
			log.Error(err.Error())
			os.Exit(1)
		}

		// Store the file name if an alternative was not provided.
		if FileCreateCmdVars.name == "" {
			FileCreateCmdVars.name = info.Name()
		}

		// Processing parents
		parents := strings.Split(FileCreateCmdVars.parents, ",")

		// Building base file.
		f := drive.File{
			Name:        FileCreateCmdVars.name,
			MimeType:    FileCreateCmdVars.mimeType,
			Parents:     parents,
			Description: FileCreateCmdVars.description,
		}

		// Building call based on args.
		driveService := mustConnectToDriveOrFail(RootVars.SACredsPath)
		call := driveService.Files.Create(&f).Media(file).SupportsAllDrives(true)
		if FileCreateCmdVars.includedLabelIDs != "" {
			call.IncludeLabels(FileCreateCmdVars.includedLabelIDs)
		}

		uploadedFile, err := call.Do()
		file.Close()
		if err != nil {
			if strings.Contains(err.Error(), "File not found: ., notFound") {
				fmt.Println("No root parent folder found. Please specific a specific parent.")
			} else {
				fmt.Printf("Failed to upload file %s\n", err.Error())
			}

			log.Error(err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", uploadedFile.Id)

	},
}

func init() {
	fileCmd.AddCommand(fileCreateCmd)

	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.filePath, "file", "", "The path to the file to upload/create.")
	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.name, "name", "", "The optional name for the file. If not provided the name of the file being uploaded will be used.")
	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.parents, "parents", "", "A comma separated list of IDs for the parent of the file. (aka folder)")
	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.description, "description", "", "A optional description to add to the file.")
	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.includedLabelIDs, "label-ids", "", "A comma separted list of label ids for labels that should be applied to the file.")
	fileCreateCmd.Flags().StringVar(&FileCreateCmdVars.description, "mimetype", "", "Will force the mimetype. If empty Google Drive attempts to determine the mime type automatically.")

	fileCreateCmd.MarkFlagRequired("file")
}
