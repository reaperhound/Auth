package main

import (
	"jwt-auth/internal/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := routes.SetupRouter()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server running on :8080")
	log.Fatal(srv.ListenAndServe())
}
