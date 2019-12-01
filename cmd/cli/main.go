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
	modules := []di.Definition{}
	modules = append(modules, di.NewCLIModules()...)
	modules = append(modules, di.NewCoreModules()...)

	c, err := di.NewContainer(modules...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.Invoke(executor(c)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executor(c *di.Container) func(app cli.App) {
	fn := func(app cli.App) {
		// set
		app.Pre(func(*cobra.Command, []string) error { return c.Invoke(register) })

		// run
		app.Run()
	}

	return fn
}

func register(targets map[string]map[string]*cobra.Command, task *handler.TaskHandler) {
	targets["task"]["new"].RunE = task.CreateTask
}
