package services

import (
	"context"

	"github.com/todo_api/models"
)

type AuthServ interface {
	CreateUser(context.Context, *models.User) (*Token, error)
	LoginUser(context.Context, *models.AuthUser) (*Token, error)
	UpdateToken(context.Context, string) (*Token, error)
}

type TaskServ interface {
	CreateTask(context.Context, *models.Task) (*models.Task, error)
	GetAllTasks(context.Context) ([]models.Task, error)
	EditTask(context.Context, *models.Task) (*models.Task, error)
	DeleteTask(context.Context, int) error
	ChangeStatus(context.Context, int, bool) (*models.Task, error)
}

type FolderServ interface {
	CreateFolder(context.Context, *models.Folder) (*models.Folder, error)
	GetAllFolders(context.Context) ([]models.Folder, error)
	DeleteFolder(context.Context, int) error
	ChangeTitle(context.Context, *models.Folder) (*models.Folder, error)
}
