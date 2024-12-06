package server

import (
	"log"
	"net/http"
	"os"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found")
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: GetRouter(),
	}

	return server
}
