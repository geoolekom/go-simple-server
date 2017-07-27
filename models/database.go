package models

type Database interface {
	SelectUser(id int) (*User, error)
	SelectLocation(id int) (*Location, error)
	SelectVisit(id int) (*Visit, error)
	InsertUser(users []User) error
	InsertLocation(locations []Location) error
	InsertVisit(visits []Visit) error
	UpdateUser(user User) error
	UpdateLocation(location Location) error
	UpdateVisit(visit Visit) error
}

type Model struct { Database }

func New(db Database) *Model {
	return &Model{
		Database: db,
	}
}
