package controllers

type Controller struct {
	BaseUrl string
	Routes map[string]string
}

func (c Controller) Url(i string) string {
	return c.BaseUrl + c.Routes[i]
}