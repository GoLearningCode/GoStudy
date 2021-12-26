package routes

import (
	"net/http"

	"github.com/GoStudy/webapp/controllers"
)

func PostRoute() {
	http.HandleFunc("/post", controllers.PostController)
}
