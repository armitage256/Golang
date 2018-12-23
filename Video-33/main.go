package main

import(
	"fmt"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {

	fmt.Println("Listening port 8080")

	templates = template.Must( template.ParseGlob("templates/*.html") )

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()

	name := keys.Get("name")

	fmt.Printf( "Seja bem vindo! %s\r\n" ,name)

	templates.ExecuteTemplate(w, "index.html", nil)

}