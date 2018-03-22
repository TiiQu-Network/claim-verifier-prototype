package controllers

import (
	"net/http"
	"github.com/gobuffalo/buffalo"
)

type MemberCertification struct {
	Controller
}

type MemberCertificationResource struct{
	buffalo.Resource
}

func (m MemberCertification) ToBlockchain(c buffalo.Context) error {
	// TODO
	return new(error)
}
