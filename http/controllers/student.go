package controllers

import (
	"net/http"
	"github.com/gobuffalo/buffalo"
)

type Student struct{
	Controller
}

type StudentResource struct{
	buffalo.Resource
}

func (s Student) ToBlockchain(c buffalo.Context) error {
	// TODO
	return new(error)
}
