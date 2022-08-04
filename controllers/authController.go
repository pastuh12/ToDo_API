package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
	"github.com/todo_api/store"
)

type AuthController struct {
	ctx     context.Context
	service *services.Manager
}

func NewAuth(ctx context.Context, store *store.Store) *AuthController {
	return &AuthController{
		ctx:     ctx,
		service: services.New(ctx, store),
	}
}

func (ctr *AuthController) Login(ctx echo.Context) error {
	var user models.AuthUser
	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode user data"))
	}
	err = ctx.Validate(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	token, err := ctr.service.Auth.LoginUser(ctx.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "This user is not registered"))
	}

	return ctx.JSON(http.StatusOK, token)
}

func (ctr *AuthController) Registration(ctx echo.Context) error {
	var user models.User
	err := ctx.Bind(&user)
	logrus.Info(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode user data"))
	}
	err = ctx.Validate(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	token, err := ctr.service.Auth.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		return errors.Wrap(err, "failed user registration")
	}

	return ctx.JSON(http.StatusOK, token)
}

func (ctr *AuthController) RenewTokens(ctx echo.Context) error {
	return nil
}
