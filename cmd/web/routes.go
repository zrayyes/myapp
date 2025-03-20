package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", app.getTasks)

	myMiddleware := chainMiddleware(app.recoverPanic, app.logRequest)

	return myMiddleware(mux)
}
