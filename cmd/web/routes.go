package main

import "net/http"


// the routes() method returns a servemux containing our application routes
func (app *application) routes() * http.ServeMux {
	mux := http.NewServeMux()

	// create a file server which serves files out of the "./ui/static" directory
	//the path provided to http.Dir is realtive to the project directory root 
	// relative to cmd 
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snipper/create", app.snippetCreate)
	return mux
}