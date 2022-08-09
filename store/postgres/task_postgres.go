package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/todo_api/models"
)

type TaskPostgres struct {
	store *PostgresDB
}

func NewTaskPostgres(s *PostgresDB) *TaskPostgres {
	return &TaskPostgres{
		store: s,
	}
}

func (r *TaskPostgres) DeleteTasksByFolderID(ctx context.Context, folderID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE folder_id = $1", tableTasks)
	err := r.store.db.QueryRow(query, folderID).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return errors.Wrap(err, "failed deletion task from folder")
	}

	return nil
}

func (r *TaskPostgres) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, folder_id) VALUES($1, $2, $3) RETURNING id", tableTasks)
	err := r.store.db.QueryRow(query, task.Title, task.Description, task.FolderID).Scan(&id)
	if err != nil {
		return nil, errors.Wrap(err, "failed create task")
	}

	task.ID = id

	return task, nil
}

func (r *TaskPostgres) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var taskList []models.Task
	var task models.Task
	query := fmt.Sprintf("SELECT id, title, description, status, folder_id FROM %s", tableTasks)
	rows, err := r.store.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return taskList, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.FolderID,
		)
		if err != nil {
			return taskList, err
		}
		taskList = append(taskList, task)
	}

	return taskList, nil
}

func (r *TaskPostgres) EditTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	query := fmt.Sprintf(
		"UPDATE %s SET title = $1, description = $2, status = $3, folder_id = $4 WHERE id = $5 RETURNING id",
		tableTasks,
	)
	err := r.store.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.FolderID,
		task.ID,
	).Scan(task.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task does not exist")
		}
		return nil, err
	}
	return task, nil
}

func (r *TaskPostgres) ChangeStatus(ctx context.Context, taskID int, status bool) (*models.Task, error) {
	var task models.Task
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2 RETURNING *", tableTasks)
	err := r.store.db.QueryRow(query, status, taskID).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.FolderID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task does not exist")
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskPostgres) DeleteTask(ctx context.Context, taskID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableTasks)
	err := r.store.db.QueryRow(query, taskID).Scan()
	if err != nil {
		return errors.Wrap(err, "failed deletion task")
	}

	return nil
}
