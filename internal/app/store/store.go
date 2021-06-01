package store

type Store interface {
	Card() CardRepository
}
