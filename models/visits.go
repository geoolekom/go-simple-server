package models

type Visit struct {
	Id int `json:"id"`
	User int `json:"user"`
	Location int `json:"location"`
	VisitedAt string `json:"visited_at"`
	Mark int `json:"mark"`
}
