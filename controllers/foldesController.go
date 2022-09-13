package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
)

type FolderController struct {
	ctx     context.Context
	service *services.Manager
}

func NewFolderController(ctx context.Context, services *services.Manager) *FolderController {
	return &FolderController{
		ctx:     ctx,
		service: services,
	}
}

func (ctr *FolderController) CreateFolder(ctx echo.Context) error {
	var folder models.Folder

	err := ctx.Bind(&folder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode user data"))
	}

	err = ctx.Validate(&folder)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	_, err = ctr.service.Folder.CreateFolder(ctx.Request().Context(), &folder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, folder)
}

func (ctr *FolderController) GetAllFolders(ctx echo.Context) error {
	folderList, err := ctr.service.Folder.GetAllFolders(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, folderList)
}

func (ctr *FolderController) DeleteFolder(ctx echo.Context) error {
	folderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	err = ctr.service.Folder.DeleteFolder(ctx.Request().Context(), folderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil

}

func (ctr *FolderController) ChangeTitle(ctx echo.Context) error {
	var folder models.Folder
	var err error
	_, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	err = ctx.Bind(&folder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode user data"))
	}

	err = ctx.Validate(&folder)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	_, err = ctr.service.Folder.ChangeTitle(ctx.Request().Context(), &folder)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, folder)
}
