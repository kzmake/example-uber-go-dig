package cli

import (
	"github.com/spf13/cobra"
)

var taskAliases = []string{"tasks", "t"}

// Task ...
var Task map[string]*cobra.Command

func init() {
	Task = NewTaskCommand()
}

// NewTaskCommand ...
func NewTaskCommand() map[string]*cobra.Command {
	cmd := map[string]*cobra.Command{
		"new":    newCommand("task", taskAliases...),
		"list":   newCommand("task", taskAliases...),
		"get":    newCommand("task -i TASK_ID", taskAliases...),
		"modify": newCommand("task -i TASK_ID", taskAliases...),
		"remove": newCommand("task -i TASK_ID", taskAliases...),
	}

	cmd["new"].PersistentFlags().StringP("data", "d", "", "set data(default: \"\")")
	cmd["get"].PersistentFlags().StringP("id", "i", "", "find by task id")
	cmd["modify"].PersistentFlags().StringP("id", "i", "", "find by task id")
	cmd["modify"].PersistentFlags().StringP("data", "d", "", "set data")
	cmd["remove"].PersistentFlags().StringP("id", "i", "", "find by task id")

	return cmd
}
