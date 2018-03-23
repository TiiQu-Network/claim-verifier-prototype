package controllers

import "github.com/gobuffalo/buffalo"

type Home struct {
	Controller
}

func (h Home) Home(c buffalo.Context) error {
	// TODO
	return new(error)
}

func (h Home) Regenerate(c buffalo.Context) error {
	// TODO
	return new(error)
}
