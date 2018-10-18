package router

import (
	"net/http"

	"../controllers"
)

func Router() {
	http.HandleFunc("/", controllers.LoginPage)
	http.HandleFunc("/index", controllers.HomePage)
}
