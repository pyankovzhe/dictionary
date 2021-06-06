package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Card struct {
	ID          int    `json:"id"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

func (c *Card) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Original, validation.Required, validation.Length(1, 1000)),
	)
}
