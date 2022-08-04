package store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/todo_api/config"
	"github.com/todo_api/store/postgres"
)

type Store struct {
	Postgres *postgres.PostgresDB

	Authtorization AuthRepo
	Task           TaskRepo
	Folder         FolderRepo
}

func New(ctx context.Context, conf *config.Config) (*Store, error) {

	//connect to Postgres
	pgDB, err := postgres.NewConnect(conf)
	if err != nil {
		return nil, errors.Wrap(err, "postgres connection failed")
	}

	//run migrations
	// if pgDB != nil {
	// 	log.Println("Running PostgreSQL migrations...")
	// 	if err := MakePostgresMigrationsUp(conf); err != nil {
	// 		return nil, errors.Wrap(err, "runPgMigrations failed")
	// 	}
	// }

	//init Store
	var store Store

	if pgDB != nil {
		store.Postgres = pgDB
		store.Authtorization = postgres.NewAuthPostgres(pgDB)
		// store.Folder = postgres.NewFolderPostgres(pgDB)
		// store.Task = postgres.NewTaskPOstgres(pgDB)
	}

	return &store, err

}
