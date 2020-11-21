package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type sqlite struct {
	*sql.DB
}

func NewDB() (*sqlite, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sqlStmt := `
	create table if not exists foo (id integer not null primary key autoincrement, name text);
	create table if not exists dummy (id integer not null primary key autoincrement, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlite{
		db,
	}, nil
}

func (s *sqlite) CloseSQLite() {
	s.Close()
}

func (s *sqlite) Insert(name string) error {
	tx, err := s.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	stmt, err := tx.Prepare("insert into foo(name) values(?)")
	if err != nil {
		return errors.WithStack(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		return errors.WithStack(err)
	}
	tx.Commit()

	return nil
}

func (s *sqlite) Select(name string) (string, error) {
	stmt, err := s.Prepare("select id, name from foo where name = ?")
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer stmt.Close()

	var (
		id          int
		fetchedName string
	)
	if err := stmt.QueryRow(name).Scan(&id, &fetchedName); err != nil {
		return "", errors.WithStack(err)
	}

	return fmt.Sprintf("ID: %d/ NAME: %s", id, fetchedName), nil
}
