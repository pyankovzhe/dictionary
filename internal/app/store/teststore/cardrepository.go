package teststore

import (
	"github.com/pyankovzhe/dictionary/internal/app/model"
	"github.com/pyankovzhe/dictionary/internal/app/store"
)

type CardRepository struct {
	store *Store
	cards map[int]*model.Card
}

func (r *CardRepository) Create(c *model.Card) error {
	if err := c.Validate(); err != nil {
		return err
	}

	c.ID = len(r.cards) + 1
	r.cards[c.ID] = c

	return nil
}

func (r *CardRepository) Find(id int) (*model.Card, error) {
	c, ok := r.cards[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return c, nil
}
