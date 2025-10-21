package main

import (
	"log"

	"github.com/Infamous003/go-blog-backend/cmd/api"
)

func main() {
	server := api.NewAPIServer(":9090", nil)

	err := server.Run()

	if err != nil {
		log.Fatal("[ERROR] Failed to start the server")
	}
}
