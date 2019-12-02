package di

import (
	"github.com/kzmake/example-uber-go-dig/infrastructure/handler"
	"github.com/kzmake/example-uber-go-dig/infrastructure/renderer"
	"github.com/kzmake/example-uber-go-dig/interface/controller"
	"github.com/kzmake/example-uber-go-dig/interface/presenter"
	"github.com/kzmake/example-uber-go-dig/usecase/interactor"
)

// NewCoreDefinitions ...
func NewCoreDefinitions() []Definition {
	defs := []Definition{
		{constructor: handler.NewTaskHandler},
		{constructor: controller.NewTaskController},
		{constructor: renderer.NewTaskRenderer},
		{constructor: interactor.NewCreateTaskInteractor},
		{constructor: presenter.NewCreateTaskPresenter},
	}

	return defs
}
