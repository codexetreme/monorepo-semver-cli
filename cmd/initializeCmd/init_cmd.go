package initializeCmd

import "github.com/spf13/cobra"

// Cmd implements utils.ICommand (subcommands also implement the interface)
type Cmd struct {
	Cmd *cobra.Command
	// declare flags here
}

func (c Cmd) GetCmd() *cobra.Command {
	return c.Cmd
}

func NewCmd() *Cmd {
	c := &Cmd{}

	c.Cmd = &cobra.Command{
		Use:   "init",
		Short: "command for initializing the cli configuration",
		//		Long: `A longer description that spans multiple lines and likely contains examples
		//and usage of using your command. For example:
		//
		//Cobra is a CLI library for Go that empowers applications.
		//This application is a tool to generate the needed files
		//to quickly create a Cobra application.`,
	}

	return c
}
