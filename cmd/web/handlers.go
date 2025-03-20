package main

import (
	"fmt"
	"net/http"
)

func (app *application) getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tasks")
}
