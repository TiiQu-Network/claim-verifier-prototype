package controllers

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

type HomeController struct {
	Controller
}

var (
	Home HomeController
)

func init() {
	Home = HomeController{Controller{
		BaseUrl: "/",
		Routes: map[string]string{
			"idx": "",
			"tut": "tutorial/",
			"reg": "regenerate/",
		},
	}}
}

func (h HomeController) Index() (string, buffalo.Handler) {
	return h.Url("idx"), func(c buffalo.Context) error {
		return c.Render(200, p.HTML("home/welcome.html"))
	}
}

func (h HomeController) Tutorial() (string, buffalo.Handler) {
	return h.Url("tut"), func(c buffalo.Context) error {
		return c.Render(200, r.HTML("home/tutorial.html"))
	}
}

func (h HomeController) Regenerate() (string, buffalo.Handler) {
	return h.Url("reg"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}
