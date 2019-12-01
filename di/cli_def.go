package di

import (
	"github.com/spf13/cobra"

	"github.com/kzmake/example-uber-go-dig/cli"
)

// NewCLIModules ...
func NewCLIModules() []Definition {
	defs := []Definition{
		// cli.App
		func(root *cobra.Command, operations []*cobra.Command, targets map[string]map[string]*cobra.Command) cli.App {
			builder := cli.New().Root(root).Operations(operations...)
			for _, target := range targets {
				builder = builder.Targets(target)
			}
			return builder.Build()
		},

		// *cobra.Command
		cli.NewRootCommand,

		// []*cobra.Command
		cli.NewOperationCommands,

		// map[string]map[string]*cobra.Command
		func() map[string]map[string]*cobra.Command {
			targets := map[string]map[string]*cobra.Command{
				"task": cli.NewTaskCommand(),
			}
			return targets
		},
	}

	return defs
}
