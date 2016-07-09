package models

import (
	"fmt"
	"time"
)

type User struct {
	Id 				int
	UserName 		string
	Password		string
	Status			string
	CreatedBy		int
	CreationDate	time.Time
}

func RetrieveUser(id int) int {
	fmt.Println("Retrieve User via Model")
	return id
}