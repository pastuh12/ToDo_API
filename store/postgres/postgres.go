package postgres

import (
	"database/sql"

	"github.com/todo_api/config"

	_ "github.com/lib/pq" //...
)

const (
	tableUsers    = "users"
	tableTasks    = "tasks"
	tableFolders  = "folders"
	tableSessions = "sessions"
)

type PostgresDB struct {
	db *sql.DB
}

func NewConnect(conf *config.Config) (*PostgresDB, error) {
	db, err := sql.Open("postgres", conf.PgURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db}, nil
}

func (s *PostgresDB) Close() {
	s.db.Close()

}
