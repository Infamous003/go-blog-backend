package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Infamous003/go-blog-backend/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
	l    *log.Logger
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	l := log.New(os.Stdout, "[API] ", log.LstdFlags|log.Lshortfile)

	return &APIServer{
		addr: addr,
		db:   db,
		l:    l,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	apiRouter := chi.NewRouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiRouter)

	r.Get("/", handleRoot)

	r.Mount("/api/v1", apiRouter)

	newServer := &http.Server{
		Addr:    s.addr,
		Handler: r,
	}

	s.l.Printf("Server running on port %s\n", s.addr)

	return newServer.ListenAndServe()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WHat's up!"))
}
