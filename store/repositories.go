package store

import (
	"context"

	"github.com/todo_api/models"
)

type AuthRepo interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, *models.AuthUser) (int, error)
	CheckSession(context.Context, *models.Session) error
	UpdateSession(context.Context, *models.Session) error
	GetSessionByToken(context.Context, string) (int, error)
}

type TaskRepo interface {
	CreateTask(context.Context, *models.Task) (*models.Task, error)
	GetAllTasks(context.Context) ([]models.Task, error)
	EditTask(context.Context, *models.Task) (*models.Task, error)
	DeleteTask(context.Context, int) error
	ChangeStatus(context.Context, int, bool) (*models.Task, error)
}

type FolderRepo interface {
	CreateFolder(context.Context, *models.Folder) (*models.Folder, error)
	GetAllFolders(context.Context) ([]models.Folder, error)
	DeleteFolder(context.Context, int) error
	EditFolder(context.Context, *models.Folder) (*models.Folder, error)
}
