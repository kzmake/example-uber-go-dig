package cli

import (
	"github.com/spf13/cobra"
)

// NewRootCommand ...
func NewRootCommand(cmds []*OperationCommand) *RootCommand {
	cmd := &RootCommand{&cobra.Command{}}
	cmd.Use = "cli"
	cmd.Short = "Controls service"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}
