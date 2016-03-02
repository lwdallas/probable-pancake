package main

// admission: shamelessley stealing an example from cory lanou, https://github.com/corylanou/tns-restful-json-api

// TODO: Authentication
// TODO: Tags?
// TODO: version control via routes

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
