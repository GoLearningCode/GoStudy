package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foobar")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/bar", bar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
