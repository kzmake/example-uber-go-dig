package renderer

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/kzmake/example-uber-go-dig/usecase/port"
)

// TaskRenderer ...
type TaskRenderer struct {
}

// NewTaskRenderer ...
func NewTaskRenderer() *TaskRenderer {
	return &TaskRenderer{}
}

// CreateTask ...
func (r *TaskRenderer) CreateTask(cmd *cobra.Command, o *port.CreateTaskOutputData) {
	cmd.Println(spew.Sdump(o))
}
