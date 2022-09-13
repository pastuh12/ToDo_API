package tests

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/evt/rest-api-example/lib/types"

	"github.com/evt/rest-api-example/lib/types"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/todo_api/controllers"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
	mock_services "github.com/todo_api/services/mocks"
	"github.com/todo_api/validator"
)

func TestController_registration(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockAuthServ)

	testUser := &models.User{
		EmailAddress: "email@gmail.com",
		Password:     "12345678",
		Name:         "Ilya",
	}

	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name:      "Ok",
			inputBody: `{"email_address":"email@gmail.com","password":"12345678","name":"Ilya"}`,
			expectations: func(ctx context.Context, s *mock_services.MockAuthServ) {
				s.EXPECT().CreateUser(gomock.All(), testUser).Return(
					&services.Token{
						AccessToken:  "1234",
						RefreshToken: "dfdf",
						ExpiresAt:    2321312,
						RefreshExt:   5434452,
					}, nil,
				)
			},
			err:  nil,
			code: http.StatusOK,
		},
		{
			name:         "without data",
			inputBody:    `{}`,
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			err:          errors.New("code=422, message=Key: 'User.EmailAddress' Error:Field validation for 'EmailAddress' failed on the 'required' tag, Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag, Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			name:         "bad request",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			inputBody:    `{some"}`,
			err:          errors.New("code=400, message=Could not decode user data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
		{
			name: "service error",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {
				svc.EXPECT().CreateUser(ctx, testUser).Return(nil, types.ErrBadRequest)
			},
			inputBody: `{ "email_address": "email@gmail.com", "password": "12345678", "name": "Ilya" }`,
			err:       errors.New("failed user registration: bad request"),
			code:      http.StatusBadRequest,
		},
	}
	for _, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		r, err := http.NewRequest(echo.POST, "/users/", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		c := gomock.NewController(t)

		svc := mock_services.NewMockAuthServ(c)
		testCase.expectations(ctx.Request().Context(), svc)

		d := controllers.NewAuth(ctx.Request().Context(), &services.Manager{Auth: svc})
		err = d.Registration(ctx)
		assert.Equal(t, testCase.err == nil, err == nil)

		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		} else {
			assert.Equal(t, testCase.code, w.Code)
		}
	}
}

func TestController_login(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockAuthServ)

	testAuthUser := &models.AuthUser{
		EmailAddress: "email@gmail.com",
		Password:     "12345678",
	}
	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name: "ok",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {
				svc.EXPECT().LoginUser(gomock.All(), testAuthUser).Return(
					&services.Token{
						AccessToken:  "1234",
						RefreshToken: "dfdf",
						ExpiresAt:    2321312,
						RefreshExt:   5434452,
					}, nil,
				)
			},
			inputBody: `{ "email_address": "email@gmail.com", "password": "12345678"}`,
			err:       nil,
			code:      http.StatusOK,
		},
		{
			name:         "without data",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			inputBody:    `{ }`,
			err:          errors.New("code=422, message=Key: 'AuthUser.EmailAddress' Error:Field validation for 'EmailAddress' failed on the 'required' tag, Key: 'AuthUser.Password' Error:Field validation for 'Password' failed on the 'required' tag"),
			code:         http.StatusBadRequest,
		},
		{
			name:         "bad request",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			inputBody:    `{some"}`,
			err:          errors.New("code=400, message=Could not decode user data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
	}

	for _, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		r, err := http.NewRequest(echo.POST, "/users/", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		c := gomock.NewController(t)

		svc := mock_services.NewMockAuthServ(c)
		testCase.expectations(ctx.Request().Context(), svc)

		d := controllers.NewAuth(ctx.Request().Context(), &services.Manager{Auth: svc})
		err = d.Login(ctx)
		assert.Equal(t, testCase.err == nil, err == nil)

		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		} else {
			assert.Equal(t, testCase.code, w.Code)
		}
	}
}

func TestController_renewTokens(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockAuthServ)
	refreshToken := "12uh32ruu3"
	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name: "ok",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {
				svc.EXPECT().UpdateToken(gomock.All(), refreshToken).Return(
					&services.Token{
						AccessToken:  "1234",
						RefreshToken: "dfdf",
						ExpiresAt:    2321312,
						RefreshExt:   5434452,
					}, nil,
				)
			},
			inputBody: `{"refreshToken": "12uh32ruu3"}`,
			err:       nil,
			code:      http.StatusOK,
		},
		{
			name:         "without data",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			inputBody:    `{ }`,
			err:          errors.New("code=422, message=Key: 'TokenRefresh.Token' Error:Field validation for 'Token' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			name:         "bad request",
			expectations: func(ctx context.Context, svc *mock_services.MockAuthServ) {},
			inputBody:    `{some"}`,
			err:          errors.New("code=400, message=Could not decode user data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
	}
	for _, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		r, err := http.NewRequest(echo.POST, "/users/", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		c := gomock.NewController(t)

		svc := mock_services.NewMockAuthServ(c)
		testCase.expectations(ctx.Request().Context(), svc)

		d := controllers.NewAuth(ctx.Request().Context(), &services.Manager{Auth: svc})
		err = d.RenewTokens(ctx)
		assert.Equal(t, testCase.err == nil, err == nil)

		if err != nil {
			if testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		} else {
			assert.Equal(t, testCase.code, w.Code)
		}
	}
}
