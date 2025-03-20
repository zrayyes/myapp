package main

import (
	"fmt"
	"net/http"
)

func (app *application) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.taskStore.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// TODO: Implement a JSON response
	fmt.Fprintf(w, "%+v", tasks)
}
