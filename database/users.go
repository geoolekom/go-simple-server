package database

import (
	"github.com/geoolekom/go-simple-server/models"
	"errors"
)

func (s Storage) SelectUser(id int) (*models.User, error) {
	stmt, err := s.connection.Prepare("SELECT id, email, first_name, last_name, gender, birth_date FROM \"user\" WHERE id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	count := 0
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Gender, &user.BirthDate)
		count ++
	}

	if count > 1 {
		return nil, errors.New("Returns more than one!")
	} else if count == 0 {
		return nil, errors.New("Not found!")
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
