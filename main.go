package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func renderHTMLTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

func HomeRender(w http.ResponseWriter, r *http.Request) {

	data := struct{}{}

	renderHTMLTemplate(w, homeTemplate, data)
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/src/").Handler(http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))

	homeTemplate = template.Must(template.ParseFiles("src/index.html"))
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/", HomeRender).Methods("GET")

	http.ListenAndServe(":4000", router)
}
