package cli

import (
	"github.com/spf13/cobra"
)

// NewOperationCommands ...
func NewOperationCommands() []*cobra.Command {
	cmds := []*cobra.Command{
		&cobra.Command{
			Use:     "new",
			Aliases: []string{"n"},
			Short:   "Create a new resource",
		},
		&cobra.Command{
			Use:     "list",
			Aliases: []string{"ls", "l"},
			Short:   "List resources",
		},
		&cobra.Command{
			Use:     "get",
			Aliases: []string{"g"},
			Short:   "Get a resource",
		},
		&cobra.Command{
			Use:     "modify",
			Aliases: []string{"mod", "m"},
			Short:   "Modify a resource",
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm", "r"},
			Short:   "Remove a resource",
		},
	}

	return cmds
}
