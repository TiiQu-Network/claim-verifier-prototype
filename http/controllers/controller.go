package controllers

import (
	"fmt"
	"github.com/TiiQu-Network/claim-verifier-prototype/app/pages"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := pages.LoadPage(title)
	if err != nil {
		p = &pages.Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}