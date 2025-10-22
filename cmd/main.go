package main

import (
	"context"
	"log"
	"time"

	"github.com/Infamous003/go-blog-backend/cmd/api"
	"github.com/Infamous003/go-blog-backend/config"
	"github.com/Infamous003/go-blog-backend/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Couldn't load .env file, falling back to defaults")
	}

	// using withtimeout cause we dont want the connection to stay alive forever
	// this ctx is good for short lived operations like db operations, io, network ops, etc
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // freeing up resources

	cfg := config.InitConfig()
	dsn := config.GetDBURL(cfg)

	dbPool, err := db.NewPostgresStorage(ctx, dsn)
	if err != nil {
		log.Fatalf("[Fatal] Failed to connect to DB: %s", err.Error())
	}
	defer dbPool.Close()

	log.Println("[DB] Connected to postgres successfully")

	server := api.NewAPIServer(":"+cfg.Port, dbPool)
	log.Printf("[API] Starting server on port %s", cfg.Port)

	if err = server.Run(); err != nil {
		log.Fatalf("[FATAL] Failed to start the server: %w", err)
	}
}
