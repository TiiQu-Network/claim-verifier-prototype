package controllers

import "github.com/gobuffalo/buffalo"

type Controller struct {
	BaseUrl string
	Routes map[string]string
	Resource buffalo.Resource
}

func (c Controller) Url(i string) string {
	if i == "" {
		return c.BaseUrl
	}
	return c.BaseUrl + c.Routes[i]
}