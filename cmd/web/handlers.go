package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"strconv"

	"snippetbox.abhijitLearning/internal/models"
)

func (app *application) home (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()

	if err != nil{
		app.serverError(w,r,err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// }

	// ParseFiles reads the files and stores the templates in a template set (ts)
	// destructuring the files slice into variablues using ...
	
	// ts, err := template.ParseFiles(files...)

	// if err != nil{
	// 	app.serverError(w,r, err)
	// 	return
	// }

	// executeTemplate() method to write the content of base

	// err = ts.ExecuteTemplate(w,"base",nil)
	// if err != nil {
	// 	app.serverError(w,r,err)
	// 	}
		
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet , err := app.snippets.Get(id)

	if err != nil{
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else {
			app.serverError(w,r,err)
		}
		return
	}


	fmt.Fprintf(w, "%+v", snippet)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed )
		return
	}
	
	title :="O snail"
	content :="O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w,r,err)
		return
	}


	http.Redirect(w,r, fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)
	w.Write([]byte("Create a new snippet..."))
}
