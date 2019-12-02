package di

import (
	"go.uber.org/dig"
)

// Container ...
type Container = dig.Container

// Definition ...
type Definition struct {
	constructor interface{}
	name        string
	group       string
}
