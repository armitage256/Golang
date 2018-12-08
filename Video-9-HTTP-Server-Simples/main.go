package main

import(
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Listening port 8080")
	
	http.HandleFunc("/", index)
	
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)

	w.Write([]byte("Hello world!"))

}