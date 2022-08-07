package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	FolderID    int    `json:"folderID"`
}
