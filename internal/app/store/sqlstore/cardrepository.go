package sqlstore

import (
	"database/sql"

	"github.com/pyankovzhe/dictionary/internal/app/model"
	"github.com/pyankovzhe/dictionary/internal/app/store"
)

type CardRepository struct {
	store *Store
}

func (r *CardRepository) Create(c *model.Card) error {
	// TODO: validation

	return r.store.db.QueryRow(
		"INSERT INTO cards (original, translation) VALUES ($1, $2) RETURNING id",
		c.Original,
		c.Translation,
	).Scan(c.ID)
}

func (r *CardRepository) Find(id int) (*model.Card, error) {
	c := *&model.Card{}

	if err := r.store.db.QueryRow(
		"SELECT id, original, translation FROM cards WHERE id = $1",
		id,
	).Scan(&c.ID, &c.Original, &c.Translation); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return &model.Card{}, nil
}
