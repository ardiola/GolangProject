package controllers

import (
	"html/template"
	"net/http"
)

type tittle struct {
	Tittle   string
	Username string
	Password string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	p := tittle{Tittle: "Home"}
	t, _ := template.ParseFiles("templates/home.html")

	t.Execute(w, p)
}
