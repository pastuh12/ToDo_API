package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
	"github.com/todo_api/store"
)

type AuthController struct {
	ctx     context.Context
	service *services.Manager
}

type TokenRefresh struct {
	Token string `json:"refreshToken"`
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
		return echo.NewHTTPError(http.StatusUnauthorized, errors.Wrap(err, "This user is not registered"))
	}

	return ctx.JSON(http.StatusOK, token)
}

func (ctr *AuthController) Registration(ctx echo.Context) error {
	var user models.User
	err := ctx.Bind(&user)
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
	var t TokenRefresh

	err := ctx.Bind(&t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "Could not decode user data"))
	}
	err = ctx.Validate(&t)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	token, err := ctr.service.Auth.UpdateToken(ctx.Request().Context(), t.Token)
	if err != nil {
		if fmt.Sprint(err) == "refreshToken not valid" {
			return echo.NewHTTPError(http.StatusUnauthorized, err) //change
		}
		if fmt.Sprint(err) == "refreshToken expired" {
			return echo.NewHTTPError(http.StatusUnauthorized, err) //change
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, token)
}
