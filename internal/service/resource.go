package service

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/kyeett/highscore-server/model"
)

func (s *Service) addScore(w http.ResponseWriter, r *http.Request) {
	var m model.Score
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.Highscore.Add(&m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) listByGame(w http.ResponseWriter, r *http.Request) {
	name, err := url.PathUnescape(chi.URLParam(r, "gameName"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := s.Highscore.ListByGame(model.Game{Name: name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&list)
}
