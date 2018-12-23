package main

import(
	"fmt"
	"net/http"
	"html/template"
	"strconv"
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

	// convertendo string para int
	n1, _ := strconv.Atoi( keys.Get("n1") ) 
	n2, _ := strconv.Atoi( keys.Get("n2") )

	var sum int = 0

	sum = (n1 + n2)

	templates.ExecuteTemplate(w, "index.html", struct{
		N1  int
		N2  int
		Sum int
	}{
		N1: n1,
		N2: n2,
		Sum: sum, 
	})

}