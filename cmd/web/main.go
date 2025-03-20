package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/zrayyes/myapp/internal/models"
)

type application struct {
	infoLog   *log.Logger
	errorLog  *log.Logger
	taskStore models.TaskStore
}

func newApplication() *application {
	logHandler := slog.NewTextHandler(os.Stdout, nil)

	return &application{
		infoLog:   slog.NewLogLogger(logHandler, slog.LevelInfo),
		errorLog:  slog.NewLogLogger(logHandler, slog.LevelError),
		taskStore: models.NewTaskStoreInMemory(),
	}
}

func main() {
	app := newApplication()

	server := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	app.infoLog.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
