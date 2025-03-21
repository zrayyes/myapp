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

func newApplication(taskStore models.TaskStore) *application {
	logHandler := slog.NewTextHandler(os.Stdout, nil)

	return &application{
		infoLog:   slog.NewLogLogger(logHandler, slog.LevelInfo),
		errorLog:  slog.NewLogLogger(logHandler, slog.LevelError),
		taskStore: taskStore,
	}
}

func main() {
	app := newApplication(models.NewTaskStoreInMemory())

	server := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	// TODO: Remove later
	app.taskStore.Create("Task 1", "This is the first task")
	app.taskStore.Create("Task 2", "This is the second task")

	app.infoLog.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
