package service

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/kyeett/highscore-server/internal/highscore"
)

type Service struct {
	Router    chi.Router
	Highscore highscore.Service
}

type orderFailure struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
	Handled   bool      `json:"handled"`
}

func New(db *sqlx.DB) *Service {

	s := &Service{
		Highscore: highscore.NewBasic(db),
	}

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
		r.Post("/", s.addScore)
		r.Route("/{gameName}", func(r chi.Router) {
			r.Get("/", s.listByGame)
		})
	})

	s.Router = r
	return s
}
