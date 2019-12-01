package interactor

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/kzmake/example-uber-go-dig/domain/entity"
	"github.com/kzmake/example-uber-go-dig/usecase/port"
)

// CreateTaskInteractor ...
type CreateTaskInteractor struct {
}

// NewCreateTaskInteractor ...
func NewCreateTaskInteractor() port.CreateTaskInputPort {
	return &CreateTaskInteractor{}
}

// Execute ...
func (i *CreateTaskInteractor) Execute(input *port.CreateTaskInputData) (*entity.Task, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	task := &entity.Task{
		ID:   id,
		Data: input.Data,
	}

	return task, nil
}
