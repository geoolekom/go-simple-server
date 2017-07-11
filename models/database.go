package models

type Database interface {
	SelectUser(id int) (*User, error)
	SelectLocation(id int) (*Location, error)
}

type Model struct { Database }

func New(db Database) *Model {
	return &Model{
		Database: db,
	}
}
