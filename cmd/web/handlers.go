package main

import (
	"net/http"
)

func (app *application) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.taskStore.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.jsonResponse(w, http.StatusOK, tasks)
}
