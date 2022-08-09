package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
	"github.com/todo_api/store"
)

type FolderController struct {
	ctx     context.Context
	service *services.Manager
}

func NewFolderController(ctx context.Context, store *store.Store) *FolderController {
	return &FolderController{
		ctx:     ctx,
		service: services.New(ctx, store),
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
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, folder)
}

func (ctr *FolderController) GetAllFolders(ctx echo.Context) error {
	folderList, err := ctr.service.Folder.GetAllFolders(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	logrus.Info(folderList)
	return ctx.JSON(http.StatusOK, folderList)
}

func (ctr *FolderController) DeleteFolder(ctx echo.Context) error {
	folderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	err = ctr.service.Folder.DeleteFolder(ctx.Request().Context(), folderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusOK, fmt.Sprintf("folder with id %d was deleted", folderID))

}

func (ctr *FolderController) ChangeTitle(ctx echo.Context) error {
	var folder models.Folder
	var err error
	folder.ID, err = strconv.Atoi(ctx.Param("id"))
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
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, folder)
}
