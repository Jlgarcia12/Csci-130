package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

var mytemp *template.Template

type userstuff struct {
	Name, Email, Message string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formhandler", formhandler)
	http.HandleFunc("/temp1", linkone)
	http.HandleFunc("/temp2", linktwo)
	http.HandleFunc("/temp3", linkthree)
	mytemp = template.Must(template.ParseFiles("form.html", "formhandler.html", "temp1.html", "temp2.html", "temp3.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "formhandler.html", struct {
		Name    string
		Email   string
		Message string
	}{r.FormValue("name"), r.FormValue("email"), r.FormValue("msg")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkone(w http.ResponseWriter, r *http.Request) {
	thisuser := r.FormValue("name") + " " + r.FormValue("email") + " " + r.FormValue("message")
	err := mytemp.ExecuteTemplate(w, "temp1.html", thisuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linktwo(w http.ResponseWriter, r *http.Request) {
	thisuser := r.FormValue("name") + " " + r.FormValue("email") + " " + r.FormValue("message")
	err := mytemp.ExecuteTemplate(w, "temp2.html", thisuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkthree(w http.ResponseWriter, r *http.Request) {
	thisuser := r.FormValue("name") + " " + r.FormValue("email") + " " + r.FormValue("message")
	err := mytemp.ExecuteTemplate(w, "temp3.html", thisuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
