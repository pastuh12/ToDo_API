package services

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/todo_api/models"
	"github.com/todo_api/store"
)

type TaskService struct {
	ctx   context.Context
	store *store.Store
}

func NewTaskService(ctx context.Context, store *store.Store) *TaskService {
	return &TaskService{
		ctx:   ctx,
		store: store,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	task, err := s.store.Task.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	taskList, err := s.store.Task.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return taskList, nil
}

func (s *TaskService) EditTask(ctx context.Context, taskID int, task *models.Task) (*models.Task, error) {
	task, err := s.store.Task.EditTask(ctx, taskID, task)
	if err != nil {
		if fmt.Sprint(err) == "no rows in result set" {
			return nil, errors.New("task does not exist")
		}
		return nil, err
	}

	return task, nil
}

func (s *TaskService) ChangeStatus(ctx context.Context, id int, newStatus bool) (*models.Task, error) {

	task, err := s.store.Task.ChangeStatus(ctx, id, newStatus)
	if err != nil {
		if fmt.Sprint(err) == "no rows in result set" {
			return nil, errors.New("task does not exist")
		}
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {

	return nil
}
