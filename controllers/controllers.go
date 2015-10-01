package controllers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "all users are here !")
}
