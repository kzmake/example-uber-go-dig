package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kzmake/example-uber-go-dig/cli"
	"github.com/kzmake/example-uber-go-dig/di"
	"github.com/kzmake/example-uber-go-dig/infrastructure/handler"
)

func main() {
	defs := []di.Definition{}
	defs = append(defs, di.NewCLIDefinitions()...)
	defs = append(defs, di.NewCoreDefinitions()...)

	c, err := di.NewContainer(defs...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.Invoke(executor(c)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executor(c *di.Container) func(app cli.AppProvider) {
	fn := func(app cli.AppProvider) {
		// set register
		app.Command().PersistentPreRunE = func(cmd *cobra.Command, args []string) error { return c.Invoke(register) }

		// run
		app.Run()
	}

	return fn
}

func register(cmds di.TaskCommands, taskHandler *handler.TaskHandler) {
	cmds.CreateTask.RunE = taskHandler.CreateTask
}
