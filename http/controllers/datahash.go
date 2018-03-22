package controllers

import (
	"net/http"
	"github.com/gobuffalo/buffalo"
)

type DataHash struct {
	Controller
}

type DataHashResource struct{
	buffalo.Resource
}
