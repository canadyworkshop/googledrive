package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var RootVars = struct {
	LoggingLevel string
	SACredsPath  string
}{
	LoggingLevel: "none",
	SACredsPath:  "",
}

var log *zap.SugaredLogger

// rootCmd represents the base command when called without subcommands.
var rootCmd = &cobra.Command{
	Use:   "googledrive",
	Short: "googledrive provides cli access to Google Drives",
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
	rootCmd.PersistentFlags().StringVar(&RootVars.LoggingLevel, "logging-level", "none", "The level to write logs at. [verbose, debug, none]")
	rootCmd.PersistentFlags().StringVarP(&RootVars.SACredsPath, "sa-credentials-file-path", "c", "", "The path to the credential file for the SA being used.")

	rootCmd.MarkPersistentFlagRequired("sa-credentials-file-path")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	switch RootVars.LoggingLevel {
	case "verbose":
		logger, _ := zap.NewProduction()
		log = logger.Sugar()
	case "debug":
		logger, _ := zap.NewDevelopment()
		log = logger.Sugar()

	default:
		log = zap.NewNop().Sugar()
	}

	log.Info("logging enabled")
}
