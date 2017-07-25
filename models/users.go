package models

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	BirthDate string `json:"birth_date"`
}
