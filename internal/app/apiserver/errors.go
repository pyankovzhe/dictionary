package apiserver

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}
