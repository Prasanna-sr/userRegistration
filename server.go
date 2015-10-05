package main

import (
	"log"
	"net/http"
	// "net/url"
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
	http.HandleFunc("/user", methodHandler)
	http.HandleFunc("/user/", fnHandler)
}

func fnHandler(response http.ResponseWriter, request *http.Request) {
	log.Print(request.URL)
	log.Print(request.RequestURI)
}

func methodHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		id := request.URL.Query().Get("emailid")
		if len(id) > 0 {
			controllers.GetUser(response, request)
		} else {
			controllers.GetAllUsers(response, request)
		}

	} else if request.Method == "POST" {
		controllers.CreateUser(response, request)
	}
}
