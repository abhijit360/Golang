package main

import (
	"net/http"
	"runtime/debug"
)

// http.StatusText(int) returns a human friendly text representation of a given HTTP status code
// e.g http.StatusText(400) returns "Bad Request"

// the serverError helper writes a log entry at Error level
// it also sense a generic 500 interrnal server error to the user

func (app *application) serverError(w http.ResponseWriter, r*http.Request, err error){
	var (
	method = r.Method
	uri = r.URL.RequestURI()
	trace = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method , "uri", uri, "trace", trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// the clientError helper sends a specific status code and corresponding description to the user
func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w,http.StatusText(status), status)
}


//convenience wrapper around client error to send a 404 when needed
func (app *application) notFound(w http.ResponseWriter){
	app.clientError(w, http.StatusNotFound)
}