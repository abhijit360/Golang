package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
)


// enables the dependencies for the application to be injected into the desired files
// build handlers against this struct
type application struct {
	logger *slog.Logger
}

func openDB(dsn string) (*sql.DB, error){
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
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

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL Data Source Name")
	// gets the addr variable from the command prompt
	flag.Parse() 
		
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	logger.Info("starting server", "addr", *addr)
	
	db, err := openDB(*dsn)
	if err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
	}

	err = http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}
