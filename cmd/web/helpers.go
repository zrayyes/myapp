package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
)

// Writes an error message and stack trace to the errorLog
// then sends a generic error 500
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Send back a specific status code and corresponding description back to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// A convenience wrapper around clientError to return 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", body)
}

func readJSON(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func getIntFromPath(r *http.Request, key string) (int, error) {
	value := r.PathValue(key)
	if value == "" {
		return 0, fmt.Errorf("missing value for key %s", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid value for key %s", key)
	}

	return intValue, nil
}
