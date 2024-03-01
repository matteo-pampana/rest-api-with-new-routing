package notes

import (
	"encoding/json"
	"net/http"
)

type NoteService interface {
	Create(note Note) Note
	ReadAll() []Note
	Read(id NoteID) (Note, bool)
	Update(id NoteID, note Note) (Note, bool)
	Delete(id NoteID) bool
}

type NoteError struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type NoteHTTPHandler struct {
	noteService NoteService
}

func NewNoteHTTPHandler(noteService NoteService) *NoteHTTPHandler {
	return &NoteHTTPHandler{
		noteService: noteService,
	}
}

func (n *NoteHTTPHandler) HandleHTTPPost(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		n.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	newNote := n.noteService.Create(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newNote)
	if err != nil {
		n.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (n *NoteHTTPHandler) HandleHTTPGet(w http.ResponseWriter, r *http.Request) {
	notes := n.noteService.ReadAll()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		n.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (n *NoteHTTPHandler) HandleHTTPGetWithID(w http.ResponseWriter, r *http.Request) {
	id := NoteID(r.PathValue("id"))
	note, found := n.noteService.Read(id)
	if !found {
		n.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(note)
	if err != nil {
		n.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (n *NoteHTTPHandler) HandleHTTPPut(w http.ResponseWriter, r *http.Request) {
	id := NoteID(r.PathValue("id"))

	var newNote Note
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		n.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	note, found := n.noteService.Update(id, newNote)
	if !found {
		n.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		n.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (n *NoteHTTPHandler) HandleHTTPDelete(w http.ResponseWriter, r *http.Request) {

	id := NoteID(r.PathValue("id"))
	found := n.noteService.Delete(id)
	if !found {
		n.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (n *NoteHTTPHandler) errorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encodingError := json.NewEncoder(w).Encode(NoteError{
		StatusCode: statusCode,
		Error:      errorString,
	})
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
	}
}
