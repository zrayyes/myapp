package main

import (
	"net/http"

	"github.com/zrayyes/myapp/internal/models"
)

func (app *application) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.taskStore.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.jsonResponse(w, http.StatusOK, tasks)
}

func (app *application) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIntFromPath(r, "id")
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	task, err := app.taskStore.Get(id)
	if err != nil {
		switch err {
		case models.ErrRecordNotFound:
			app.notFound(w)
		default:
			app.serverError(w, err)
		}
		return
	}

	app.jsonResponse(w, http.StatusOK, task)
}
