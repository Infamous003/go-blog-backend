package main

import (
	"context"
	"log"
	"time"

	"github.com/Infamous003/go-blog-backend/config"
	"github.com/Infamous003/go-blog-backend/internal/server"
	"github.com/Infamous003/go-blog-backend/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Couldn't load .env file, falling back to defaults")
	}

	// using withtimeout cause we dont want the connection to stay alive forever
	// this dbCtx is good for short lived operations like db operations, io, network ops, etc
	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // freeing up resources

	cfg := config.InitConfig()

	dbPool, err := db.NewPostgresStorage(dbCtx, config.GetDBURL(cfg))
	if err != nil {
		log.Fatalf("[ERROR] failed to connect to DB: %v", err)
	}
	defer dbPool.Close()

	log.Printf("[INFO] Connection to postgres established")

	appServer := server.New(&cfg, dbPool)
	log.Printf("[API] Starting server on port %s", cfg.Port)

	if err := appServer.Run(); err != nil {
		log.Fatalf("[FATAL] Server error: %v", err)
	}
}
