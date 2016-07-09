package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

import helper "github.com/jespersei/webserver/helpers"
import model "github.com/jespersei/webserver/models"


func UserRetrieval(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Retrieve User via Model from the Controller.")
	id, _ := strconv.Atoi(helper.GetParam("userId", r))
	model.RetrieveUser(id)
}