package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Fuck %s", name)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server started on :8090")
	http.ListenAndServe(":8090", nil)
}