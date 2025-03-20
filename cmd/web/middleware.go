package main

import (
	"fmt"
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always run in the event of a panic as Go unwinds the stack)
		defer func() {
			// Use recover() to check if a panic happened or not
			if err := recover(); err != nil {
				// Tell Go's HTTP server to automatically close the current connection after the response has been sent.
				w.Header().Set("Connection", "close")
				// We format the error with Errorf as it has type any (string, error, or something else)
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func chainMiddleware(mw ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for i := len(mw) - 1; i >= 0; i-- {
			next = mw[i](next)
		}
		return next
	}
}
