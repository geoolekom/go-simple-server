package database

import (
	"github.com/geoolekom/go-simple-server/models"
	"errors"
)

func (s Storage) SelectVisit(id int) (*models.Visit, error) {
	rows, err := s.visitSelector.Query(id)
	if err != nil {
		return nil, err
	}
	var visit models.Visit
	count := 0
	for rows.Next() {
		err = rows.Scan(&visit.Id, &visit.User, &visit.Location, &visit.VisitedAt, &visit.Mark)
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
	return &visit, nil
}

func (s Storage) InsertVisit(visits []models.Visit) error {
	for _, visit := range visits {
		_, err := s.visitInsert.Exec(visit.Id, visit.User, visit.Location, visit.VisitedAt, visit.Mark)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Storage) UpdateVisit(visit models.Visit) error {
	rows, err := s.visitUpdate.Exec(visit.Id, visit.User, visit.Location, visit.VisitedAt, visit.Mark)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("404")
	}
	return err
}
