package controllers

import (
	"github.com/gobuffalo/buffalo"
)

type Institution struct {
	Controller
}

type InstitutionResource struct {
	buffalo.Resource
}

func (i Institution) Students(c buffalo.Context) error {
	// TODO
	return new(error)
}

func (i Institution) ToBlockchain(c buffalo.Context) error {
	// TODO
	return new(error)
}

func (i Institution) Blockchain(c buffalo.Context) error {
	// TODO
	return new(error)
}
