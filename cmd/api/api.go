package api

import (
	"database/sql"
	"net/http"

	"github.com/Infamous003/go-blog-backend/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	apiRouter := chi.NewRouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiRouter)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("What's up!"))
	})

	r.Mount("/api/v1", apiRouter)

	newServer := &http.Server{
		Addr:    s.addr,
		Handler: r,
	}
	return newServer.ListenAndServe()
}
