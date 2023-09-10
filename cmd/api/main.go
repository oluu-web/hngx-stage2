package main

import (
	"log"
	"net/http"

	"github.com/oluu-web/hngx-stage2/cmd/api/models"
	"github.com/oluu-web/hngx-stage2/cmd/api/routes"
)

func main() {
	err := models.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %w", err)
	}

	router := routes.InitRoutes()
	port := "4000"
	log.Printf("Server is listening of port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
