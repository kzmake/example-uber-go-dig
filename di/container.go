package di

import (
	"fmt"

	"go.uber.org/dig"
)

// NewContainer ...
func NewContainer(defs ...Definition) (*Container, error) {
	container := dig.New()

	for _, d := range defs {
		if err := container.Provide(d.constructor, dig.Name(d.name), dig.Group(d.group)); err != nil {
			return nil, fmt.Errorf("failed injection: %v", err)
		}
	}

	return container, nil
}
