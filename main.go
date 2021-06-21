package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
	"monorepo-semver-cli/cmd"
	"monorepo-semver-cli/error_roster"
	"os"
)


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
