package main

import (
	"log"
	"net/http"
	"user-registration/controllers"
)

func main() {
	port := "8080"
	routeHandler()
	log.Print("Server running in port " + port)
	http.ListenAndServe(":"+port, nil)
}

func routeHandler() {
	http.HandleFunc("/user", userController.GetAllUsers)
}
