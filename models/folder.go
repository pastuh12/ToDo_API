package models

type Folder struct {
	ID    int    `json:"id"`
	Title string `json:"title" validate:"required"`
}
