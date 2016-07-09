package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Server has started.")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}