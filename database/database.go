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
	s.userSelector, err = s.connection.Prepare("SELECT id, email, first_name, last_name, gender, birth_date FROM \"user\" WHERE id=$1")
	return err
}

func (s *Storage) Close() {
	s.userSelector.Close()
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

