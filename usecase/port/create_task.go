package port

import (
	"github.com/kzmake/example-uber-go-dig/domain/entity"
	"github.com/kzmake/example-uber-go-dig/domain/vo"
)

// CreateTaskInputData ...
type CreateTaskInputData struct {
	Data vo.TaskData
}

// CreateTaskOutputData ...
type CreateTaskOutputData struct {
	Task *entity.Task
}

// CreateTaskInputPort ...
type CreateTaskInputPort interface {
	Execute(*CreateTaskInputData) (*entity.Task, error)
}

// CreateTaskOutputPort ...
type CreateTaskOutputPort interface {
	Execute(*entity.Task) *CreateTaskOutputData
}
