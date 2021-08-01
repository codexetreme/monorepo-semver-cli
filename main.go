package main

import (
    "os"

    "github.com/codexetreme/monorepo-semver-cli/cmd"
    "github.com/codexetreme/monorepo-semver-cli/error_roster"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra/doc"
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
    DrawBox()
    DrawList()
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

