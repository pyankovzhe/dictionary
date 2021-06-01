package model

type Card struct {
	ID          int    `json:"id"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}
