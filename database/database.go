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
	locationSelector *sql.Stmt
	locationInsert	 *sql.Stmt
	visitSelector    *sql.Stmt
	visitInsert	     *sql.Stmt
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
	s.userSelector, err = s.connection.Prepare("SELECT id, email, first_name, last_name, gender, birth_date FROM \"user\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.userInsert, err = s.connection.Prepare("INSERT INTO \"user\" (id, email, first_name, last_name, gender, birth_date) VALUES ($1, $2, $3, $4, $5, to_date($6, 'DD.MM.YYYY'))")
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
	s.visitSelector, err = s.connection.Prepare("SELECT id, user, location, visited_at, mark FROM \"visit\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.visitInsert, err = s.connection.Prepare("INSERT INTO \"visit\" (id, \"user\", location, visited_at, mark) VALUES ($1, $2, $3, to_timestamp($4, 'DD.MM.YYYY HH24:MI:SS'), $5)")
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() {
	s.userSelector.Close()
	s.userInsert.Close()
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

