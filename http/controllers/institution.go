package controllers

import (
	"github.com/gobuffalo/buffalo"
	"errors"
)

type InstitutionController struct {
	Controller
}

type InstitutionResource struct {
	buffalo.Resource
}

var Institution InstitutionController

func init(){
	Institution = InstitutionController{
		Controller{
			BaseUrl:"/institution/",
			Routes: map[string]string{
				"stu":"{institution}/students/",
				"2bc":"{institution}/students/toBlockchain/",
				"bcn":"{institution}/students/blockchain/",
			},
			Resource:InstitutionResource{},
		},
	}
}

func (i InstitutionController) Students() (string, buffalo.Handler) {
	return i.Url("stu"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}

func (i InstitutionController) ToBlockchain() (string, buffalo.Handler) {
	return i.Url("2bc"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}

func (i InstitutionController) Blockchain() (string, buffalo.Handler) {
	return i.Url("bcn"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}
