package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"path/filepath"
)

type Storage struct {
	connection       *sql.DB
	userSelector     *sql.Stmt
	userInsert       *sql.Stmt
	userUpdate       *sql.Stmt
	locationSelector *sql.Stmt
	locationInsert	 *sql.Stmt
	locationUpdate   *sql.Stmt
	visitSelector    *sql.Stmt
	visitInsert	     *sql.Stmt
	visitUpdate      *sql.Stmt
}

func (s *Storage) createTablesIfNotExist() error {
	ddlFile, _ := filepath.Abs("sql/ddl.sql")
	dat, err := ioutil.ReadFile(ddlFile)
	if err != nil {
		return err
	}
	createSql := string(dat)
	_, err = s.connection.Exec(createSql)
	return err
}

func (s *Storage) prepareStatements() (err error) {
	s.userSelector, err = s.connection.Prepare("SELECT id, email, first_name, last_name, gender, to_char(birth_date, 'DD.MM.YYYY') AS birth_date FROM \"user\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.userInsert, err = s.connection.Prepare("INSERT INTO \"user\" (id, email, first_name, last_name, gender, birth_date) VALUES ($1, $2, $3, $4, $5, to_date($6, 'DD.MM.YYYY'))")
	if err != nil {
		return err
	}
	s.userUpdate, err = s.connection.Prepare("UPDATE \"user\" SET id = $1, email = $2, first_name = $3, last_name = $4, gender = $5, birth_date = to_date($6, 'DD.MM.YYYY') WHERE id = $1")
	if err != nil {
		return err
	}
	s.locationSelector, err = s.connection.Prepare("SELECT id, place, country, city, distance FROM \"location\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.locationInsert, err = s.connection.Prepare("INSERT INTO \"location\" (id, place, country, city, distance) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	s.locationUpdate, err = s.connection.Prepare("UPDATE \"location\" SET id = $1, place = $2, country = $3, city = $4, distance = $5 WHERE id = $1")
	if err != nil {
		return err
	}
	s.visitSelector, err = s.connection.Prepare("SELECT id, \"user\", location, to_char(visited_at, 'DD.MM.YYYY HH24:MI:SS') AS visited_at, mark FROM \"visit\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.visitInsert, err = s.connection.Prepare("INSERT INTO \"visit\" (id, \"user\", location, visited_at, mark) VALUES ($1, $2, $3, to_timestamp($4, 'DD.MM.YYYY HH24:MI:SS'), $5)")
	if err != nil {
		return err
	}
	s.visitUpdate, err = s.connection.Prepare("UPDATE \"visit\" SET id = $1, \"user\" = $2, location = $3, visited_at = to_timestamp($4, 'DD.MM.YYYY HH24:MI:SS'), mark = $5 WHERE id = $1")
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() {
	s.userSelector.Close()
	s.userInsert.Close()
	s.userUpdate.Close()
	s.locationSelector.Close()
	s.locationInsert.Close()
	s.visitSelector.Close()
	s.visitInsert.Close()
	s.connection.Close()
}

func InitDatabase(initString string) (*Storage, error) {
	if connection, err := sql.Open("postgres", initString); err != nil {
		return nil, err
	} else {
		db := &Storage{connection: connection}
		if err := db.connection.Ping(); err != nil {
			return nil, err
		}
		if err := db.createTablesIfNotExist(); err != nil {
			return nil, err
		}
		if err := db.prepareStatements(); err != nil {
			return nil, err
		}
		return db, nil
	}
}

