package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", app.getTasks)
	mux.HandleFunc("GET /tasks/{id}", app.getTask)
	mux.HandleFunc("POST /tasks", app.createTask)

	myMiddleware := chainMiddleware(app.recoverPanic, app.logRequest)

	return myMiddleware(mux)
}
