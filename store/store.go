package store

import (
	"database/sql"

	"github.com/todo_api/config"

	_ "github.com/lib/pq" //...
)

type Store struct {
	config *config.Config
	db     *sql.DB
}

func New(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.PgURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()

}
