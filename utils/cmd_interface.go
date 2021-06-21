package utils

import "github.com/spf13/cobra"

type ICommand interface {
	GetCmd() *cobra.Command
}
