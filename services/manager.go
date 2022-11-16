package services

import (
	"context"

	"github.com/todo_api/store"
)

type Manager struct {
	Auth   AuthServ
	Task   TaskServ
	Folder FolderServ
}

func New(ctx context.Context, store *store.Store) *Manager {
	return &Manager{
		Auth:   NewAuthService(ctx, store),
		Task:   NewTaskService(ctx, store),
		Folder: NewFolderService(ctx, store),
	}
}
