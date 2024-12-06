package main

import (
	"log"

	"github.com/Evilcmd/data-metric-framework/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	server := server.NewServer()

	log.Println("Starting server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error starting the server")
	}
}
