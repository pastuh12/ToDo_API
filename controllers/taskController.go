package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
)

type TaskController struct {
	ctx     context.Context
	service *services.Manager
}

type Status struct {
	St bool `json:"status" validate:"required"`
}

func NewTask(ctx context.Context, services *services.Manager) *TaskController {
	return &TaskController{
		ctx:     ctx,
		service: services,
	}
}

func (ctr *TaskController) AddTask(ctx echo.Context) error {
	var task models.Task
	err := ctx.Bind(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode task data"))
	}
	err = ctx.Validate(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	logrus.Info(task)

	_, err = ctr.service.Task.CreateTask(ctx.Request().Context(), &task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, task)
}

func (ctr *TaskController) GetAllTasks(ctx echo.Context) error {

	taskList, err := ctr.service.Task.GetAllTasks(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, taskList)
}

func (ctr *TaskController) EditTask(ctx echo.Context) error {
	var task models.Task
	var err error
	task.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad param not exist")
	}

	err = ctx.Bind(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode task data"))
	}

	err = ctx.Validate(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	_, err = ctr.service.Task.EditTask(ctx.Request().Context(), &task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "service error"))
	}

	return ctx.JSON(http.StatusOK, task)
}

func (ctr *TaskController) ChangeStatus(ctx echo.Context) error {
	var status Status

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = ctx.Bind(&status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode status"))
	}

	err = ctx.Validate(&status)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	task, err := ctr.service.Task.ChangeStatus(ctx.Request().Context(), id, status.St)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, task)
}

func (ctr *TaskController) DeleteTask(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	err = ctr.service.Task.DeleteTask(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "service error"))
	}

	return nil
}
