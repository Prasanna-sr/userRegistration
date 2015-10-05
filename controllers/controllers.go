package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-registration/models"
)

type status struct {
	Message string
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users := models.FetchAllUsers()
	m, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(response, string(m))
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	emailId := request.URL.Query().Get("emailid")
	user := models.FetchUser(emailId)
	m, err := json.Marshal(user)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(response, string(m))
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	type userData struct {
		EmailID  string
		Password string
		Name     string
		City     string
	}
	ud := userData{}
	err := json.NewDecoder(request.Body).Decode(&ud)
	if err != nil {
		log.Print(err)
	}
	userErr := models.CreateUser(ud.EmailID, ud.Password, ud.Name, ud.City)
	s := status{}
	if userErr != nil {
		s.Message = "Invalid input"
		message, err := json.Marshal(s)
		if err != nil {
			log.Print(err)
		}
		http.Error(response, string(message), 400)
	} else {
		s.Message = "Success"
		message, err := json.Marshal(s)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(response, string(message))
	}
}
