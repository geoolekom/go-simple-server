package database

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	"path/filepath"
)

type Storage struct {
	connection *sql.DB
	userSelector *sql.Stmt
	locationSelector *sql.Stmt
	visitSelector *sql.Stmt
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
	s.locationSelector, err = s.connection.Prepare("SELECT id, place, country, city, distance FROM \"location\" WHERE id = $1")
	if err != nil {
		return err
	}
	s.visitSelector, err = s.connection.Prepare("SELECT id, user, location, visited_at, mark FROM \"visit\" WHERE id = $1")
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() {
	s.userSelector.Close()
	s.locationSelector.Close()
	s.visitSelector.Close()
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

