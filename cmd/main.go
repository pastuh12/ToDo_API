package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/todo_api/config"
	"github.com/todo_api/controllers"
	"github.com/todo_api/store"
	"github.com/todo_api/validator"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	conf := config.Get()

	// l := logger.Get()

	store, err := store.New(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "failed connect to database")
	}

	taskController := controllers.NewTask(ctx, store)
	authController := controllers.NewAuth(ctx, store)
	folderController := controllers.NewFolderController(ctx, store)

	e := echo.New()
	e.Validator = validator.NewValidator()
	// Disable Echo JSON logger in debug mode
	if conf.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api")
	// User routes
	taskRoutes := v1.Group("/task")
	taskRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.SigningKey),
	}))
	taskRoutes.POST("/", taskController.AddTask)
	taskRoutes.GET("/", taskController.GetAllTasks)
	taskRoutes.DELETE("/:id", taskController.DeleteTask)
	taskRoutes.PATCH("/:id", taskController.ChangeStatus)
	taskRoutes.PUT("/:id", taskController.EditTask)

	folderRoutes := v1.Group("/folder")
	folderRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.SigningKey),
	}))
	folderRoutes.GET("/", folderController.GetAllFolders)
	folderRoutes.POST("/", folderController.CreateFolder)
	folderRoutes.DELETE("/:id", folderController.DeleteFolder)
	folderRoutes.PATCH("/:id", folderController.ChangeTitle)

	v1.POST("/registration", authController.Registration)
	v1.POST("/login", authController.Login)

	e.POST("/user/token-renew", authController.RenewTokens)

	s := &http.Server{
		Addr:         conf.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
