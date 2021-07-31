package sqlstore

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pyankovzhe/dictionary/internal/app/store"
)

type Store struct {
	db       *sql.DB
	cardRepo *CardRepository
}

// TODO: maybe not the best place for db initializing
func NewDB(driverName string, databaseURL string, ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open(driverName, databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
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
