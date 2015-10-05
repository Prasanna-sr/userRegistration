package main

import (
	"log"
	"net/http"
	"user-registration/controllers"
	"user-registration/models"
)

func main() {
	port := "8080"
	http.HandleFunc("/user", methodHandler)
	models.Connect()
	log.Print("Server running in port " + port)
	http.ListenAndServe(":"+port, nil)
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
