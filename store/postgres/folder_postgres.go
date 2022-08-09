package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/todo_api/models"
)

type FolderPostgres struct {
	store *PostgresDB
}

func NewFolderPostgres(store *PostgresDB) *FolderPostgres {
	return &FolderPostgres{
		store: store,
	}
}

func (r *FolderPostgres) CreateFolder(ctx context.Context, folder *models.Folder) (*models.Folder, error) {
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES($1) RETURNING id", tableFolders)
	err := r.store.db.QueryRow(query, folder.Title).Scan(&folder.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed creation folder")
	}

	return folder, nil
}

func (r *FolderPostgres) GetAllFolders(ctx context.Context) ([]models.Folder, error) {
	var folderList []models.Folder
	var folder models.Folder

	query := fmt.Sprintf("SELECT * FROM %s", tableFolders)
	rows, err := r.store.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return folderList, nil
		}
		return nil, errors.Wrap(err, "failed getting folders")
	}

	for rows.Next() {
		err = rows.Scan(
			&folder.ID,
			&folder.Title,
		)
		if err != nil {
			return folderList, err
		}
		folderList = append(folderList, folder)
	}

	return folderList, nil
}

func (r *FolderPostgres) DeleteFolder(ctx context.Context, folderID int) error {
	err := NewTaskPostgres(r.store).DeleteTasksByFolderID(ctx, folderID)
	if err != nil {

		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", tableFolders)
	err = r.store.db.QueryRow(query, folderID).Scan(&folderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.Wrap(err, "failed deletion folder")
		}
		return err
	}

	return nil
}

func (r *FolderPostgres) EditFolder(ctx context.Context, folder *models.Folder) (*models.Folder, error) {
	logrus.Info(folder.ID)

	query := fmt.Sprintf("UPDATE %s SET title = $1 WHERE id = $2 RETURNING title", tableFolders)
	err := r.store.db.QueryRow(query, folder.Title, folder.ID).Scan(&folder.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Folder does not exist")
		}
		return nil, err
	}

	return folder, nil
}
