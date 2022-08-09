package services

import (
	"context"

	"github.com/todo_api/models"
	"github.com/todo_api/store"
)

type FolderService struct {
	ctx   context.Context
	store *store.Store
}

func NewFolderService(ctx context.Context, store *store.Store) *FolderService {
	return &FolderService{
		ctx:   ctx,
		store: store,
	}
}

func (s *FolderService) CreateFolder(ctx context.Context, folder *models.Folder) (*models.Folder, error) {
	folder, err := s.store.Folder.CreateFolder(ctx, folder)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (s *FolderService) GetAllFolders(ctx context.Context) ([]models.Folder, error) {
	folderList, err := s.store.Folder.GetAllFolders(ctx)
	if err != nil {
		return nil, err
	}

	return folderList, nil
}

func (s *FolderService) DeleteFolder(ctx context.Context, folderID int) error {
	err := s.store.Folder.DeleteFolder(ctx, folderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *FolderService) ChangeTitle(ctx context.Context, folder *models.Folder) (*models.Folder, error) {
	folder, err := s.store.Folder.EditFolder(ctx, folder)
	if err != nil {
		return nil, err
	}

	return folder, nil
}
