package cli

import (
	"github.com/spf13/cobra"
)

var taskAliases = []string{"tasks", "t"}

// NewCreateTaskCommand ...
func NewCreateTaskCommand() *CreateContentCommand {
	cmd := &CreateContentCommand{&cobra.Command{}}
	cmd.Use = "task"
	cmd.Aliases = taskAliases
	cmd.Short = "Create a new task"
	cmd.RunE = unimplemented()

	cmd.PersistentFlags().StringP("data", "d", "", "set data(default: \"\")")

	return cmd
}

// NewListTaskCommand ...
func NewListTaskCommand() *ListContentCommand {
	cmd := &ListContentCommand{&cobra.Command{}}
	cmd.Use = "task"
	cmd.Aliases = taskAliases
	cmd.Short = "List tasks"
	cmd.RunE = unimplemented()

	return cmd
}

// NewGetTaskCommand ...
func NewGetTaskCommand() *GetContentCommand {
	cmd := &GetContentCommand{&cobra.Command{}}
	cmd.Use = "task -i TASK_ID"
	cmd.Aliases = taskAliases
	cmd.Short = "get a task"
	cmd.RunE = unimplemented()

	cmd.PersistentFlags().StringP("id", "i", "", "find by task id")

	return cmd
}

// NewModifyTaskCommand ...
func NewModifyTaskCommand() *ModifyContentCommand {
	cmd := &ModifyContentCommand{&cobra.Command{}}
	cmd.Use = "task -i TASK_ID"
	cmd.Aliases = taskAliases
	cmd.Short = "modify a task"
	cmd.RunE = unimplemented()

	cmd.PersistentFlags().StringP("id", "i", "", "find by task id")
	cmd.PersistentFlags().StringP("data", "d", "", "set data")

	return cmd
}

// NewRemoveTaskCommand ...
func NewRemoveTaskCommand() *RemoveContentCommand {
	cmd := &RemoveContentCommand{&cobra.Command{}}
	cmd.Use = "task -i TASK_ID"
	cmd.Aliases = taskAliases
	cmd.Short = "remove a task"
	cmd.RunE = unimplemented()

	cmd.PersistentFlags().StringP("id", "i", "", "find by task id")

	return cmd
}
