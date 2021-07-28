package model

import (
	"github.com/google/uuid"
)

type Account struct {
	ID    uuid.UUID `json:"id"`
	Login string    `json:"login"`
}
