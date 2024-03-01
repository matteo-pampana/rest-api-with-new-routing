package main

import (
	"net/http"

	"github.com/matteo-pampana/rest-api-with-new-routing/notes"
)

func main() {
	mux := http.NewServeMux()

	noteService := notes.NewService()
	noteManager := notes.NewNoteHTTPHandler(noteService)

	mux.HandleFunc("POST /notes", noteManager.HandleHTTPPost)
	mux.HandleFunc("GET /notes", noteManager.HandleHTTPGet)
	mux.HandleFunc("GET /notes/{id}", noteManager.HandleHTTPGetWithID)
	mux.HandleFunc("PUT /notes/{id}", noteManager.HandleHTTPPut)
	mux.HandleFunc("DELETE /notes/{id}", noteManager.HandleHTTPDelete)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
