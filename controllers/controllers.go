package controllers

import (
	"fmt"
	"log"
	"net/http"
	"user-registration/models"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "all users are here !")
	users := models.FetchAllUsers()
	log.Print(users)
}
