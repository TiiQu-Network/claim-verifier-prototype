package controllers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

type HomeController struct {
	Controller
}

var Home *HomeController

func init(){
	Home = new(HomeController)
}

func (h HomeController) Index(c buffalo.Context) error {
	return c.Render(200, p.HTML("home/welcome.html"))
}

func (h HomeController) Tutorial(c buffalo.Context) error {
	return c.Render(200, r.HTML("home/tutorial.html"))
}

func (h HomeController) Regenerate(c buffalo.Context) error {
	// TODO
	return errors.New("")
}
