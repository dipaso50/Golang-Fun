package handler

import (
	"Go_notes/server/notes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetAll retrieve all notes
func GetAll(currentWarehouse notes.NoteWarehouse) func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		resp := currentWarehouse.GetAll()
		respondJSON(w, 200, resp)
	}

	return f

}

//Hello say hello word
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

//GetByID return a notes with the noteid
func GetByID(currentWarehouse notes.NoteWarehouse) func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID := vars["noteId"]
		resp := currentWarehouse.GetById(noteID)
		respondJSON(w, 200, resp)
	}

	return f
}

//Delete a note with noteid
func Delete(currentWarehouse notes.NoteWarehouse) func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID := vars["noteId"]
		currentWarehouse.Delete(noteID)
		respondJSON(w, 200, "Deleted "+noteID)
	}

	return f
}

//Add add a note
func Add(currentWarehouse notes.NoteWarehouse) func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		var n notes.Note
		json.NewDecoder(r.Body).Decode(&n)

		fmt.Println(n)

		n2 := notes.NewNote(n.Msg)

		currentWarehouse.Add(n2)
	}

	return f
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
