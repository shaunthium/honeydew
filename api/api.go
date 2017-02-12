package api

import (
	"log"
	"net/http"
)

// This package modified from https://github.com/corylanou/tns-restful-json-api/

func RunServer() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
