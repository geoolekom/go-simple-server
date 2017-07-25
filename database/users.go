package database

import (
	"github.com/geoolekom/go-simple-server/models"
	"errors"
)

func (s Storage) SelectUser(id int) (*models.User, error) {
	rows, err := s.userSelector.Query(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	count := 0
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Gender, &user.BirthDate)
		if err != nil {
			return nil, err
		}
		count ++
	}

	if count > 1 {
		return nil, errors.New("500")
	} else if count == 0 {
		return nil, errors.New("404")
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s Storage) InsertUser(user *models.User) error {
	_, err := s.userInsert.Exec(user.Id, user.Email, user.FirstName, user.LastName, user.Gender, user.BirthDate)
	return err
}
