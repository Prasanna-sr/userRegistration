package main

import (
	"log"
	"net/http"
	"user-registration/controllers"
	"user-registration/models"
)

func main() {
	port := "8080"
	routeHandler()
	models.Connect()
	log.Print("Server running in port " + port)
	http.ListenAndServe(":"+port, nil)
}

func routeHandler() {
	http.HandleFunc("/user", controllers.GetAllUsers)
}
