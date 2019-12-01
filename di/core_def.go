package di

import (
	"github.com/kzmake/example-uber-go-dig/infrastructure/handler"
	"github.com/kzmake/example-uber-go-dig/infrastructure/renderer"
	"github.com/kzmake/example-uber-go-dig/interface/controller"
	"github.com/kzmake/example-uber-go-dig/interface/presenter"
	"github.com/kzmake/example-uber-go-dig/usecase/interactor"
)

// NewCoreModules ...
func NewCoreModules() []Definition {
	defs := []Definition{

		handler.NewTaskHandler,
		controller.NewTaskController,
		renderer.NewTaskRenderer,
		interactor.NewCreateTaskInteractor,
		presenter.NewCreateTaskPresenter,
	}

	return defs
}
