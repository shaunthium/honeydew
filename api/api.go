package api

import (
	"log"
	"net/http"
)

func RunServer() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
