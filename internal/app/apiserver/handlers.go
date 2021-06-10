package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pyankovzhe/dictionary/internal/app/model"
)

type cardRequest struct {
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

type cardResponse struct {
	*model.Card
}

func (res *cardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *server) CreateCard(w http.ResponseWriter, r *http.Request) {
	req := &cardRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	card := &model.Card{
		Original:    req.Original,
		Translation: req.Translation,
	}

	if err := s.store.Card().Create(card); err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &cardResponse{Card: card})
}
