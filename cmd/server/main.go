package main

import (
	"email-alert-api/api"
	"email-alert-api/config"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv() // Load configurations
	router := api.NewRouter()

	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
