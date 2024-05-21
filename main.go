package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)

	mux.HandleFunc("/view", snippetView)
	mux.HandleFunc("/create", snippetCreate)

	log.Print("Starting server at port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
