package store

import "github.com/pyankovzhe/dictionary/internal/app/model"

type CardRepository interface {
	Create(*model.Card) error
	Find(int) (*model.Card, error)
}
