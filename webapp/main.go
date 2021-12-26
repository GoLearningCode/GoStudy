package main

import (
	"log"
	"net/http"

	"github.com/GoStudy/webapp/routes"
)

func main() {
	routes.RootRoute()
	routes.PostRoute()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
