package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

const contentTypeHeader string = "Content-Type"
const deleteMethod string = "DELETE"
const getMethod string = "GET"
const htmlContentType string = "text/html"
const indexHtml string = "html/index.html"
const static string = "./static/"
const usersPath string = "/users"

var tmpl *template.Template

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir(static))
	r.Handle(static, http.StripPrefix(static, fs))

	tmpl = template.Must(template.ParseFiles(path.Join(static, indexHtml)))

	if tmpl == nil {
		log.Fatal("No template")
	}

	userRooter := r.PathPrefix(usersPath).Subrouter()

	userRooter.HandleFunc("", AllUsers)
	userRooter.HandleFunc("/{userId}", GetUser).Methods(getMethod)
	userRooter.HandleFunc("/{userId}", DeleteUser).Methods(deleteMethod)

	log.Printf("About to listen on 8000. Go to https://127.0.0.1:8000/")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeHeader, htmlContentType)

	msg := "All users endpoint " + r.URL.Path
	err := tmpl.Execute(w, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeHeader, htmlContentType)

	vars := mux.Vars(r)
	userId := vars["userId"]

	msg := "You have requested the user " + userId
	err := tmpl.Execute(w, msg)
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeHeader, htmlContentType)

	vars := mux.Vars(r)
	userId := vars["userId"]

	msg := "You have requested to delete the user " + userId
	err := tmpl.Execute(w, msg)
	if err != nil {
		log.Fatal(err)
	}
}
