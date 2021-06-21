package main

import (
	"github.com/codexetreme/monorepo-semver-cli/cmd"
	"github.com/codexetreme/monorepo-semver-cli/error_roster"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
	"os"
)

type loggingOptions struct {
	outputLogsInJson bool
	logLevel         string
}

type semverOptions struct {
	prefixWithV bool
}

type options struct {
}

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Info("A walrus appears")
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")
	setupCobra()
}

func setupCobra() {
	rootCmd := cmd.NewRootCmd().Cmd
	if os.Getenv("MSC_CLI_BUILD_DOCS") == "1" {
		err := doc.GenMarkdownTree(rootCmd, "docs")
		if err != nil {
			log.Fatal(err)
		}
	}
	error_roster.CheckErr(rootCmd.Execute())
}
