package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {

	fmt.Println("Listening port 8080")

	templates = template.Must(template.ParseGlob("templates/*.html"))

	router := mux.NewRouter()

	router.HandleFunc("/", index)

	fileServer := http.FileServer(http.Dir("./assets"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", nil)

}