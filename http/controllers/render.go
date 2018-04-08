package controllers

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var p *render.Engine
var assetsBox = packr.NewBox("../../public")
var templateBox = packr.NewBox("../../templates")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "layouts/application.html",

		// Box containing all of the templates:
		TemplatesBox: templateBox,
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
		},
	})

	p = render.New(render.Options{
		TemplatesBox: templateBox,
		AssetsBox:    assetsBox,
	})
}
