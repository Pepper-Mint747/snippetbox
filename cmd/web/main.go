package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// fileServer serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/".
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server is listening on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
