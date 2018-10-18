package main

import (
	"log"
	"net/http"

	//"./models"
	"./router"
	"./shared/database"
)

func main() {

	//connect database
	database.ConnectDB()
	//end

	//call route for html
	router.Router()
	//end

	//models.HomeTes()

	//call css and js for html
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	//end

	log.Fatal(http.ListenAndServe(":9000", nil))
}
