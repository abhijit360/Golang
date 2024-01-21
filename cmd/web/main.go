package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)


// enables the dependencies for the application to be injected into the desired files
// build handlers against this struct
type application struct {
	logger *slog.Logger
}

func main() {

	// type config struct {
	// 	addr string
	// 	staticDir string
	// }

	// var cfg config
	// //  to store the values into an existing variable using pointers
	// flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP Network Address")
	// flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "path to static assets")
	
	// define a command-line flag named addr with a default value
	// flag will be stored in the addr variable at runtime
	// the message will show when you add the "-help" flag when calling go run
	addr := flag.String("addr", ":4000", "HTPP network Address")

	// gets the addr variable from the command prompt
	flag.Parse() 
		
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	// create a file server which serves files out of the "./ui/static" directory
	//the path provided to http.Dir is realtive to the project directory root 
	// relative to cmd 
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))


	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	app.logger.Info("starting server on", "addr",*addr)
	err := http.ListenAndServe(*addr, mux)
	app.logger.Error(err.Error())
	os.Exit(1)
}
