package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// ParseFiles reads the files and stores the templates in a template set (ts)
	// destructuring the files slice into variablues using ...
	ts, err := template.ParseFiles(files...)

	if err != nil{
		app.serverError(w,r, err)
		return
	}

	// executeTemplate() method to write the content of base
	err = ts.ExecuteTemplate(w,"base",nil)
	if err != nil {
		app.serverError(w,r,err)
		}
		
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed )
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
