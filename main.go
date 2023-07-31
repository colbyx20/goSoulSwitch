package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func renderHTMLTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tmpl.Execute(w, data)
}

func HomeRender(w http.ResponseWriter, r *http.Request) {

	data := struct{}{}

	renderHTMLTemplate(w, homeTemplate, data)
}

func EnterEmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse Data", http.StatusInternalServerError)
	}

	email := r.FormValue("mlist")
	fmt.Println(email)

	data := struct{}{}

	renderHTMLTemplate(w, homeTemplate, data)
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/src/").Handler(http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))

	homeTemplate = template.Must(template.ParseFiles("src/index.html"))
	router.Use(mux.CORSMethodMiddleware(router))

	router.Use(loggingMiddleware)
	router.HandleFunc("/", HomeRender).Methods("GET")
	router.HandleFunc("/emailSignUp", EnterEmail)

	http.ListenAndServe(":4000", router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
