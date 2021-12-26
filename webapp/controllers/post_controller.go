package controllers

import (
	"fmt"
	"net/http"
)

func PostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "welcome to post method!")
}
