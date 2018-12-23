package main

import(
	"fmt"
	"net/http"
	"html/template"
	"strings"
)

var templates *template.Template

func main() {

	fmt.Println("Listening port 8080")

	templates = template.Must( template.ParseGlob("templates/*.html") )

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	name := r.PostForm.Get("name")

	name = strings.TrimSpace(name)

	fmt.Printf( "Seja bem vindo! %s\r\n" ,name)

	templates.ExecuteTemplate(w, "index.html", struct{
		Name string
	}{
		Name: name,
	})

}