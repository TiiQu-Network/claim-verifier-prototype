package controllers

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

type RecipientController struct {
	Controller
}

type RecipientResource struct {
	buffalo.Resource
}

var Recipient RecipientController

func init(){
	Recipient = RecipientController{
		Controller{
			BaseUrl:"/recipient/",
			Routes: map[string]string{
				"2bc":"{recipient}/to-blockchain/",
			},
			Resource: RecipientResource{},
		},
	}
}

func (r RecipientController) ToBlockchain() (string, buffalo.Handler) {
	return r.Url("2bc"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}
