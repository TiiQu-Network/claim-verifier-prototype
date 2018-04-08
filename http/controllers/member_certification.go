package controllers

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

type MemberCertificationController struct {
	Controller
	Resource buffalo.Resource
}

type MemberCertificationResource struct {
	buffalo.Resource
}

var MemberCertification MemberCertificationController

func init() {
	MemberCertification = MemberCertificationController{
		Controller{
			BaseUrl: "/member-certification/",
			Routes: map[string]string{
				"2bc": "{memberCertification}/to-blockchain/",
			},
		},
		InstitutionResource{},
	}
}

func (m MemberCertificationController) ToBlockchain() (string, buffalo.Handler) {
	return m.Url("2bc"), func(c buffalo.Context) error {
		// TODO
		return errors.New("")
	}
}
