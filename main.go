package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Shows struct {
	Location string
	Venue    string
	Date     string
}

var MailingList []string
var ShowsList = []Shows{
	{"Swindon UK", "Level III", "08/18/23"},
	{"Cambridge", "The Six Six Bar", "08/19/23"},
	{"Bridgewater", "Cobblestones", "08/22/23"},
	{"Bournemouth", "Anvil RockBar", "08/23/23"},
	{"Bristol", "The Gryphon", "08/24/23"},
	{"Brimingham", "Subside Bar", "08/25/23"},
	{"Leicestershire", "the Victoria Bikers Pub", "08/26/23"},
	{"Sheffield", "02 Academy Sheffield (H&HG Only)", "08/27/23"},
	{"Manchester", "Rebellion", "08/29/23"},
	{"Llandudno", "The Motorsports Lounge", "08/30/23"},
	{"Cardiff", "Fuel Rock Club", "09/01/23"},
	{"Kent", "Sandwich Rock Club", "09/02/23"},
}

var homeTemplate *template.Template

func renderHTMLTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tmpl.Execute(w, data)
}

func HomeRender(w http.ResponseWriter, r *http.Request) {
	renderHTMLTemplate(w, homeTemplate, ShowsList)
}

func EnterEmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse Data", http.StatusInternalServerError)
	}

	email := r.FormValue("mlist")
	MailingList = append(MailingList, email)

	for _, val := range MailingList {
		fmt.Printf("%v\n", val)
	}

	renderHTMLTemplate(w, homeTemplate, ShowsList)
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
