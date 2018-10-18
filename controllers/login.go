package controllers

import (
	"net/http"
	"text/template"
)

type tes struct {
	Tittle string
	//Hello  string
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	p := tes{Tittle: "Login Page"}
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(w, p)

}
