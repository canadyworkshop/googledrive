package cmd

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"os"
)

// mustConnectToDriveOrFail will attempt to connect to GoogleDrive using the creds provided. If
// the connection fails for any reason an error will be reported and os.Exit(1) will be executed.
func mustConnectToDriveOrFail(credsJsonFilePath string) *drive.Service {
	log.Debugf("connecting to google drive using path %s", credsJsonFilePath)
	driveService, err := drive.NewService(context.Background(), option.WithCredentialsFile(credsJsonFilePath))
	if err != nil {
		log.Error(err.Error())
		fmt.Printf("failed to connect to drive: %s", err.Error())
		os.Exit(1)
	}

	return driveService
}
