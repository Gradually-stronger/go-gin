package ctl

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	container.Provide(NewDemo)
	container.Provide(NewLogin)
	container.Provide(NewUser)
	return nil
}
