package models

type Location struct {
	Id int `json:"id"`
	Place string `json:"place"`
	Country string `json:"country"`
	City string `json:"city"`
	Distance int `json:"distance"`
}
