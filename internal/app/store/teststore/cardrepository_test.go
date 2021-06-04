package teststore_test

import (
	"testing"

	"github.com/pyankovzhe/dictionary/internal/app/store"
	"github.com/pyankovzhe/dictionary/internal/app/store/teststore"
)

func TestCardRepository_Create(t *testing.T) {
	s := teststore.New()
	c := model.TestCard(t)
	assert.NoError(t, s.Card().Create(c))
	assert.NotNil(t, c.ID)
}

func TestCardRepository_Find(t *testing.T) {
	s := teststore.New()
	_, err := s.Card().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	c := model.TestCard(t)
	s.Account().Create(c)
	c, err = s.Card().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
