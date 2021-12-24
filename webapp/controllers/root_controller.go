package controllers

import (
	"fmt"
	"net/http"
)

func RootController(writer http.ResponseWriter, response *http.Request) {
	fmt.Fprintf(writer, "Welcome to the web app!")
}
