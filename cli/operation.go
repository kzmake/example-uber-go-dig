package cli

import (
	"github.com/spf13/cobra"
)

// NewCreateCommand ...
func NewCreateCommand(cmds []*CreateContentCommand) *OperationCommand {
	cmd := &OperationCommand{&cobra.Command{}}
	cmd.Use = "new"
	cmd.Aliases = []string{"n"}
	cmd.Short = "Create a new content"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}

// NewListCommand ...
func NewListCommand(cmds []*ListContentCommand) *OperationCommand {
	cmd := &OperationCommand{&cobra.Command{}}
	cmd.Use = "list"
	cmd.Aliases = []string{"ls", "l"}
	cmd.Short = "List resources"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}

// NewGetCommand ...
func NewGetCommand(cmds []*GetContentCommand) *OperationCommand {
	cmd := &OperationCommand{&cobra.Command{}}
	cmd.Use = "get"
	cmd.Aliases = []string{"g"}
	cmd.Short = "Get a resource"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}

// NewModifyCommand ...
func NewModifyCommand(cmds []*ModifyContentCommand) *OperationCommand {
	cmd := &OperationCommand{&cobra.Command{}}
	cmd.Use = "modify"
	cmd.Aliases = []string{"mod", "m"}
	cmd.Short = "Modify a resource"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}

// NewRemoveCommand ...
func NewRemoveCommand(cmds []*RemoveContentCommand) *OperationCommand {
	cmd := &OperationCommand{&cobra.Command{}}
	cmd.Use = "remove"
	cmd.Aliases = []string{"rm", "r"}
	cmd.Short = "Remove a resource"

	for _, c := range cmds {
		cmd.AddCommand(c.Command)
	}

	return cmd
}
