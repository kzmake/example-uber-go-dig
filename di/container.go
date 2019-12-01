package di

import (
	"fmt"

	"go.uber.org/dig"
)

// NewContainer ...
func NewContainer(defs ...Definition) (*Container, error) {
	container := dig.New()

	for _, d := range defs {
		if err := container.Provide(d); err != nil {
			return nil, fmt.Errorf("di error: %v", err)
		}
	}

	return container, nil
}
