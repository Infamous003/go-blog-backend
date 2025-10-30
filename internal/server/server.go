package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Infamous003/go-blog-backend/config"
	"github.com/Infamous003/go-blog-backend/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg *config.Config
	db  *pgxpool.Pool
	l   *log.Logger
}

func New(cfg *config.Config, db *pgxpool.Pool) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
		l:   log.New(os.Stdout, "[API] ", log.LstdFlags|log.Lshortfile),
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	apiRouter := chi.NewMux()
	r.Use(middleware.Logger)

	userRepo := user.NewRepository(s.db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(apiRouter)

	r.Mount("/api/v1", apiRouter)

	httpServer := &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: r,
	}

	s.l.Printf("Server listening on %s\n", s.cfg.Port)
	return httpServer.ListenAndServe()
}
