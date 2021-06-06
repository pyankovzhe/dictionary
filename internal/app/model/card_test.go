package model_test

import (
	"testing"

	"github.com/pyankovzhe/dictionary/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestCard_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		card    func() *model.Card
		isValid bool
	}{
		{
			name: "valid",
			card: func() *model.Card {
				return model.TestCard(t)
			},
			isValid: true,
		},
		{
			name: "empty original",
			card: func() *model.Card {
				c := model.TestCard(t)
				c.Original = ""

				return c
			},
			isValid: false,
		},
		{
			name: "empty translation",
			card: func() *model.Card {
				c := model.TestCard(t)
				c.Translation = ""

				return c
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.card().Validate())
			} else {
				assert.Error(t, tc.card().Validate())
			}
		})
	}
}
