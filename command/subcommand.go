package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

// PipelineCommand is a Command implementation for specifying the pipeline.
type PipelineCommand struct {
	UI cli.Ui
}

//Help returns the help for the command
func (c *PipelineCommand) Help() string {
	helpText := `
Usage: [app name] [subcommand] [options]

  A description of the application

Options:

  -key=val                     description of the key
`
	return strings.TrimSpace(helpText)
}

//Run runs the command
func (c *PipelineCommand) Run(args []string) int {

	var hclFile string
	cmdFlags := flag.NewFlagSet("subcommand", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	cmdFlags.StringVar(&hclFile, "key", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	//c.UI.Output(args[0])
	return 0
}

//Synopsis resturns the synopsis for the command
func (c *PipelineCommand) Synopsis() string {
	return "Synopsis of [subcommand]"
}
