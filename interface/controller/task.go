package controller

import (
	"github.com/kzmake/example-uber-go-dig/usecase/port"
)

// TaskController ...
type TaskController struct {
	createIn  port.CreateTaskInputPort
	createOut port.CreateTaskOutputPort
}

// NewTaskController ...
func NewTaskController(input port.CreateTaskInputPort, output port.CreateTaskOutputPort) *TaskController {
	return &TaskController{
		createIn:  input,
		createOut: output,
	}
}

// CreateTask ...
func (c *TaskController) CreateTask(data string) *port.CreateTaskOutputData {
	in := &port.CreateTaskInputData{Data: data}
	_t, _ := c.createIn.Execute(in)

	return c.createOut.Execute(_t)
}
