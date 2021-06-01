package sqlstore

import (
	"database/sql"

	"github.com/pyankovzhe/dictionary/internal/app/store"
)

type Store struct {
	db       *sql.DB
	cardRepo *CardRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Card() store.CardRepository {
	if s.cardRepo != nil {
		return s.cardRepo
	}

	s.cardRepo = &CardRepository{
		store: s,
	}

	return s.cardRepo
}
