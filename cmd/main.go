package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/todo_api/config"
	"github.com/todo_api/store"
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
	// ctx := context.Background()
	conf := config.Get()

	// l := logger.Get()

	store := store.New(conf)
	if err := store.Open(); err != nil {
		return errors.Wrap(err, "Не удалось подключиться к базе данных")
	}

	e := echo.New()
	// Disable Echo JSON logger in debug mode
	if conf.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	s := &http.Server{
		Addr:         conf.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
