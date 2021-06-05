package model

import "testing"

func TestCard(t *testing.T) *Card {
	return &Card{
		Original:    "original word",
		Translation: "translation",
	}
}
