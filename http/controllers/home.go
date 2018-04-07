package controllers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

type Home struct {
	Controller
}

func (h Home) Home(c buffalo.Context) error {
	// TODO
	return errors.New("")
}

func (h Home) Regenerate(c buffalo.Context) error {
	// TODO
	return errors.New("")
}
