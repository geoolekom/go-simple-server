package database

import (
	"github.com/geoolekom/go-simple-server/models"
	"errors"
)

func (s Storage) SelectLocation(id int) (*models.Location, error) {
	rows, err := s.locationSelector.Query(id)
	if err != nil {
		return nil, err
	}
	var location models.Location
	count := 0
	for rows.Next() {
		err = rows.Scan(&location.Id, &location.Place, &location.Country, &location.City, &location.Distance)
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
	return &location, nil
}

func (s Storage) InsertLocation(locations []models.Location) error {
	for _, location := range locations {
		_, err := s.locationInsert.Exec(location.Id, location.Place, location.Country, location.City, location.Distance)
		if err != nil {
			return err
		}
	}
	return nil
}