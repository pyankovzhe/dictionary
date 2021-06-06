package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pyankovzhe/dictionary/internal/app/model"
)

type cardResp struct {
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

func (s *server) CreateCard(w http.ResponseWriter, r *http.Request) {
	req := &cardResp{}

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
}
