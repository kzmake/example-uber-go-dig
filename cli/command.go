package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// HandlerFn binds `func(*cobra.Command, []string) error`
type HandlerFn = func(*cobra.Command, []string) error

// ResourceCommand defines .
type ResourceCommand struct {
	new    *cobra.Command
	list   *cobra.Command
	get    *cobra.Command
	modify *cobra.Command
	remove *cobra.Command
}

func newCommand(usage string, aliases ...string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     usage,
		Aliases: aliases,
		RunE:    func(*cobra.Command, []string) error { return fmt.Errorf("unimplement") },
	}
	cmd.Short = fmt.Sprintf("Operate %s", cmd.Name())

	return cmd
}
