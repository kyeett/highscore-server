package service

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Service struct {
	chi.Router
	db ScoreDB
}

type orderFailure struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
	Handled   bool      `json:"handled"`
}

func New() *Service {

	// Create a new data base

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	r.Mount("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Route("/highscore", func(r chi.Router) {
		// r.Get("/", listOrderFailures)
		// r.Post("/mark_as_handled", markAsHandled)
		// r.Post("/mark_as_unhandled", markAsUnhandled)
	})

	return &Service{
		Router: r,
		db:     NewScoreDB(),
	}
}
