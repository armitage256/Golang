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

	html := `<select> <option value='1'> Jane Doe </option> <option value='2'> Jon Doe </option> </select>`

	templates.ExecuteTemplate(w, "index.html", struct{
		Tag template.HTML
	}{
		Tag: template.HTML( html ),
	})

}




