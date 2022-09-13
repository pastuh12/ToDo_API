package tests

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/evt/rest-api-example/lib/types"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/todo_api/config"
	"github.com/todo_api/controllers"
	"github.com/todo_api/models"
	"github.com/todo_api/services"
	mock_services "github.com/todo_api/services/mocks"
	"github.com/todo_api/validator"
)

func TestController_addTask(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockTaskServ)

	testTask := &models.Task{
		Title:       "Task",
		Status:      false,
		Description: "test task",
	}

	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name:      "ok",
			inputBody: `{"title": "Task", "status": false, "description": "test task"}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().CreateTask(gomock.All(), testTask).Return(
					testTask, nil,
				)
			},
			err:  nil,
			code: http.StatusCreated,
		},
		{
			name:         "doesn't exist required parameter",
			inputBody:    `{"status": false}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New("code=422, message=Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			name:         "without parameters",
			inputBody:    `{}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New("code=422, message=Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			name:         "bad request",
			inputBody:    `{some"}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New("code=400, message=Could not decode task data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
		{
			name:      "service error",
			inputBody: `{"title": "Task", "status": false, "description": "test task"}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().CreateTask(gomock.All(), testTask).Return(nil, types.ErrBadRequest)
			},
			err:  errors.New("code=400, message=bad request"),
			code: http.StatusBadRequest,
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

		svc := mock_services.NewMockTaskServ(c)
		testCase.expectations(ctx.Request().Context(), svc)
		d := controllers.NewTask(ctx.Request().Context(), &services.Manager{Task: svc})
		err = d.AddTask(ctx)
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

func TestController_getAllTask(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockTaskServ)
	testTable := []struct {
		name         string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name: "ok",
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().GetAllTasks(gomock.All()).Return(
					[]models.Task{}, nil,
				)
			},
			err:  nil,
			code: http.StatusOK,
		},
		{
			name: "service error",
			expectations: func(ctx context.Context, svc *mock_services.MockTaskServ) {
				svc.EXPECT().GetAllTasks(ctx).Return(nil, types.ErrBadRequest)
			},
			err:  errors.New("code=400, message=bad request"),
			code: http.StatusBadRequest,
		},
	}

	for _, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		group := e.Group("/")
		group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.Get().SigningKey),
		}))
		r, err := http.NewRequest(echo.GET, "/", strings.NewReader(""))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		c := gomock.NewController(t)

		svc := mock_services.NewMockTaskServ(c)
		testCase.expectations(ctx.Request().Context(), svc)
		d := controllers.NewTask(ctx.Request().Context(), &services.Manager{Task: svc})
		err = d.GetAllTasks(ctx)
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

func TestController_editTask(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockTaskServ)

	testTask := &models.Task{
		ID:          1,
		Title:       "Task",
		Status:      false,
		Description: "test task",
	}

	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name:      "ok",
			inputBody: `{"title": "Task", "status": false, "description": "test task"}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().EditTask(gomock.All(), testTask).Return(testTask, nil)
			},
			err:  nil,
			code: http.StatusOK,
		},
		{
			name:         "without path param id",
			inputBody:    `{"title": "Task", "status": false, "description": "test task"}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New("code=400, message=bad param not exist"),
			code:         http.StatusBadRequest,
		},
		{
			name:         "without data",
			inputBody:    `{}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New("code=422, message=Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"),
			code:         http.StatusBadRequest,
		},
		{
			name:         "bad request",
			expectations: func(ctx context.Context, svc *mock_services.MockTaskServ) {},
			inputBody:    `{some"}`,
			err:          errors.New("code=400, message=Could not decode task data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
		{
			name: "service error",
			expectations: func(ctx context.Context, svc *mock_services.MockTaskServ) {
				svc.EXPECT().EditTask(ctx, testTask).Return(nil, types.ErrBadRequest)
			},
			inputBody: `{"title": "Task", "status": false, "description": "test task"}`,
			err:       errors.New("code=400, message=service error: bad request"),
			code:      http.StatusBadRequest,
		},
	}

	for i, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		group := e.Group("/")
		group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.Get().SigningKey),
		}))
		r, err := http.NewRequest(echo.GET, "/:id", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		w := httptest.NewRecorder()

		ctx := e.NewContext(r, w)
		if i != 1 {
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
		}

		c := gomock.NewController(t)
		svc := mock_services.NewMockTaskServ(c)
		testCase.expectations(ctx.Request().Context(), svc)
		d := controllers.NewTask(ctx.Request().Context(), &services.Manager{Task: svc})
		err = d.EditTask(ctx)
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

func TestController_changeStatus(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockTaskServ)

	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name:      "ok",
			inputBody: `{}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().DeleteTask(gomock.All(), 1).Return(nil)
			},
			err:  nil,
			code: http.StatusOK,
		},
		{
			name:         "without path param id",
			inputBody:    `{}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New(`code=404, message=strconv.Atoi: parsing "": invalid syntax`),
			code:         http.StatusBadRequest,
		},
	}

	for i, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		group := e.Group("/")
		group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.Get().SigningKey),
		}))
		r, err := http.NewRequest(echo.GET, "/:id", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		w := httptest.NewRecorder()

		ctx := e.NewContext(r, w)
		if i != 1 {
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
		}

		c := gomock.NewController(t)
		t.Logf("i: %v", i)
		svc := mock_services.NewMockTaskServ(c)
		testCase.expectations(ctx.Request().Context(), svc)
		d := controllers.NewTask(ctx.Request().Context(), &services.Manager{Task: svc})
		err = d.ChangeStatus(ctx)
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

func TestController_deleteTask(t *testing.T) {
	type MockBehavior func(ctx context.Context, s *mock_services.MockTaskServ)

	testTable := []struct {
		name         string
		inputBody    string
		expectations MockBehavior
		err          error
		code         int
	}{
		{
			name:      "ok",
			inputBody: `{"status": true}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {
				s.EXPECT().DeleteTask(gomock.All(), 1).Return(nil)
			},
			err:  nil,
			code: http.StatusOK,
		},
		{
			name:         "without path param id",
			inputBody:    `{"status": true}`,
			expectations: func(ctx context.Context, s *mock_services.MockTaskServ) {},
			err:          errors.New(`code=404, message=strconv.Atoi: parsing "": invalid syntax`),
			code:         http.StatusBadRequest,
		},
		{
			name: "service error",
			expectations: func(ctx context.Context, svc *mock_services.MockTaskServ) {
				svc.EXPECT().DeleteTask(ctx, 1).Return(types.ErrBadRequest)
			},
			inputBody: `{}`,
			err:       errors.New("code=400, message=service error: bad request"),
			code:      http.StatusBadRequest,
		},
	}

	for i, testCase := range testTable {
		t.Logf("running %v", testCase.name)

		e := echo.New()
		e.Validator = validator.NewValidator()
		group := e.Group("/")
		group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.Get().SigningKey),
		}))
		r, err := http.NewRequest(echo.GET, "/:id", strings.NewReader(testCase.inputBody))
		if err != nil {
			t.Fatal("could not create request ", err)
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		w := httptest.NewRecorder()

		ctx := e.NewContext(r, w)
		if i != 1 {
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
		}

		c := gomock.NewController(t)
		t.Logf("i: %v", i)
		svc := mock_services.NewMockTaskServ(c)
		testCase.expectations(ctx.Request().Context(), svc)
		d := controllers.NewTask(ctx.Request().Context(), &services.Manager{Task: svc})
		err = d.DeleteTask(ctx)
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
