package teststore

import (
	"github.com/pyankovzhe/dictionary/internal/app/model"
	"github.com/pyankovzhe/dictionary/internal/app/store"
)

type Store struct {
	cardRepo *CardRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Card() store.CardRepository {
	if s.cardRepo != nil {
		return s.cardRepo
	}

	s.cardRepo = &CardRepository{
		store: s,
		cards: make(map[int]*model.Card),
	}

	return s.cardRepo
}
