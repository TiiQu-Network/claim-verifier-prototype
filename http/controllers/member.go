package controllers

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

type MemberController struct {
	Controller
	Resource buffalo.Resource
}

type MemberResource struct {
	buffalo.Resource
}

var Member MemberController

func init(){
	Member = MemberController{
		Controller{
			BaseUrl:"/member/",
			Routes: map[string]string{
				"crts":"{member}/certifications",
				"add":"{member}/certifications/add/",
				"bcn":"{member}/blockchain/",
			},
		},
		MemberResource{},
	}
}

func (m MemberController) Certifications() (string, buffalo.Handler) {
	return m.Url("crts"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}

func (m MemberController) AddCertification() (string, buffalo.Handler) {
	return m.Url("add"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}

func (m MemberController) Blockchain() (string, buffalo.Handler) {
	return m.Url("bcn"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}
