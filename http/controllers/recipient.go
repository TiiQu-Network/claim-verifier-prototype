package controllers

import (
	"errors"
	"github.com/gobuffalo/buffalo"
)

type Recipient struct {
	Controller
}

type RecipientResource struct {
	buffalo.Resource
}

func (r Recipient) ToBlockchain(c buffalo.Context) error {
	// TODO
	return errors.New("")
}
