package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-registration/models"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users := models.FetchAllUsers()
	m, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(response, string(m))
}
