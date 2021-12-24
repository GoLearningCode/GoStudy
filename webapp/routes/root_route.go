package routes

import (
	"net/http"

	"github.com/GoStudy/webapp/controllers"
)

func RootRoute() {
	http.HandleFunc("/", controllers.RootController)
}
