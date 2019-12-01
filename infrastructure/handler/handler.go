package handler

import (
	"github.com/spf13/cobra"

	"github.com/kzmake/example-uber-go-dig/infrastructure/renderer"
	"github.com/kzmake/example-uber-go-dig/interface/controller"
)

// TaskHandler ...
type TaskHandler struct {
	controller *controller.TaskController
	renderer   *renderer.TaskRenderer
}

// NewTaskHandler ...
func NewTaskHandler(
	controller *controller.TaskController,
	renderer *renderer.TaskRenderer,
) *TaskHandler {
	return &TaskHandler{controller: controller, renderer: renderer}
}

// CreateTask ...
func (h *TaskHandler) CreateTask(cmd *cobra.Command, args []string) error {
	data, err := cmd.PersistentFlags().GetString("data")
	if err != nil {
		return err
	}

	o := h.controller.CreateTask(data)
	h.renderer.CreateTask(cmd, o)

	return nil
}
