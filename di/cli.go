package di

import (
	"go.uber.org/dig"

	"github.com/kzmake/example-uber-go-dig/cli"
)

// TaskCommands ...
type TaskCommands struct {
	dig.In
	CreateTask *cli.CreateContentCommand `name:"create_task"`
	ListTask   *cli.ListContentCommand   `name:"list_task"`
	GetTask    *cli.GetContentCommand    `name:"get_task"`
	ModifyTask *cli.ModifyContentCommand `name:"modify_task"`
	RemoveTask *cli.RemoveContentCommand `name:"remove_task"`
}

type opretationParams struct {
	dig.In
	Cmds []*cli.OperationCommand `group:"operation"`
}
type createParams struct {
	dig.In
	Cmds []*cli.CreateContentCommand `group:"create"`
}
type listParams struct {
	dig.In
	Cmds []*cli.ListContentCommand `group:"list"`
}
type getParams struct {
	dig.In
	Cmds []*cli.GetContentCommand `group:"get"`
}
type modifyParams struct {
	dig.In
	Cmds []*cli.ModifyContentCommand `group:"modify"`
}
type removeParams struct {
	dig.In
	Cmds []*cli.RemoveContentCommand `group:"remove"`
}

// NewCLIDefinitions ...
func NewCLIDefinitions() []Definition {
	defs := []Definition{
		// app
		{constructor: cli.New},

		// root
		{constructor: cli.NewRootCommand},

		// operations
		{constructor: cli.NewCreateCommand, group: "operation"},
		{constructor: cli.NewListCommand, group: "operation"},
		{constructor: cli.NewGetCommand, group: "operation"},
		{constructor: cli.NewModifyCommand, group: "operation"},
		{constructor: cli.NewRemoveCommand, group: "operation"},

		// contents
		{constructor: cli.NewCreateTaskCommand, name: "create_task"},
		{constructor: cli.NewListTaskCommand, name: "list_task"},
		{constructor: cli.NewGetTaskCommand, name: "get_task"},
		{constructor: cli.NewModifyTaskCommand, name: "modify_task"},
		{constructor: cli.NewRemoveTaskCommand, name: "remove_task"},

		// group
		{constructor: func(p TaskCommands) *cli.CreateContentCommand { return p.CreateTask }, group: "create"},
		{constructor: func(p TaskCommands) *cli.ListContentCommand { return p.ListTask }, group: "list"},
		{constructor: func(p TaskCommands) *cli.GetContentCommand { return p.GetTask }, group: "get"},
		{constructor: func(p TaskCommands) *cli.ModifyContentCommand { return p.ModifyTask }, group: "modify"},
		{constructor: func(p TaskCommands) *cli.RemoveContentCommand { return p.RemoveTask }, group: "remove"},

		// support
		{constructor: func(p opretationParams) []*cli.OperationCommand { return p.Cmds }},
		{constructor: func(p createParams) []*cli.CreateContentCommand { return p.Cmds }},
		{constructor: func(p listParams) []*cli.ListContentCommand { return p.Cmds }},
		{constructor: func(p getParams) []*cli.GetContentCommand { return p.Cmds }},
		{constructor: func(p modifyParams) []*cli.ModifyContentCommand { return p.Cmds }},
		{constructor: func(p removeParams) []*cli.RemoveContentCommand { return p.Cmds }},
	}

	return defs
}
