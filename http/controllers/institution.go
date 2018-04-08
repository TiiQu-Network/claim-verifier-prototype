package controllers

import (
	"github.com/gobuffalo/buffalo"
	"errors"
)

type InstitutionController struct {
	Controller
	Resource buffalo.Resource
}

type InstitutionResource struct {
	buffalo.Resource
}

var Institution InstitutionController

func init(){
	Institution = InstitutionController{
		Controller{BaseUrl:"/institution/"},
		InstitutionResource{},
	}
}

func (i InstitutionController) Students(c buffalo.Context) error {
	// TODO
	return errors.New("")
}

func (i InstitutionController) ToBlockchain(c buffalo.Context) error {
	// TODO
	return errors.New("")
}

func (i InstitutionController) Blockchain(c buffalo.Context) error {
	// TODO
	return errors.New("")
}
