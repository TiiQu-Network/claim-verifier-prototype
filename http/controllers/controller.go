package controllers

type Controller struct {
	BaseUrl string
	Routes map[string]string
}

func (c Controller) Url(p string) string {
	return c.BaseUrl + p
}