package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//application struct holds application-wide dependencies for the web application
type application struct {
	errorLog *log.Logger
	infolog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infolog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server is listening on port %s...", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
