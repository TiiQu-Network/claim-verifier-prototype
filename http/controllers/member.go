package controllers

import (
	"github.com/gobuffalo/buffalo"
)

type Member struct {
	Controller
}

type MemberResource struct {
	buffalo.Resource
}

func (m Member) Certifications(c buffalo.Context) error {
	// TODO
	return new(error)
}

func (m Member) AddCertification(c buffalo.Context) error {
	// TODO
	return new(error)
}

func (m Member) Blockchain(c buffalo.Context) error {
	// TODO
	return new(error)
}
