package main

import (
	"Go_notes/server/handler"
	"Go_notes/server/memorywarehouse"
	"Go_notes/server/notes"
	"net/http"

	"github.com/gorilla/mux"
)

var currentWarehouse notes.NoteWarehouse

const (
	//WEBSERVERPORT port where the application listen
	WEBSERVERPORT = ":8080"
)

func main() {
	r := mux.NewRouter()

	currentWarehouse = memorywarehouse.NewMemoryWarehouse()

	n := notes.NewNote("Este es un test")
	currentWarehouse.Add(n)

	n2 := notes.NewNote("Esto es otro test")
	currentWarehouse.Add(n2)

	r.HandleFunc("/", handler.Hello).Methods("GET")
	r.HandleFunc("/add", handler.Add(currentWarehouse)).Methods("POST")
	r.HandleFunc("/getAll", handler.GetAll(currentWarehouse)).Methods("GET")
	r.HandleFunc("/getById/{noteId}", handler.GetByID(currentWarehouse)).Methods("GET")
	r.HandleFunc("/delete/{noteId}", handler.Delete(currentWarehouse)).Methods("DELETE")

	http.ListenAndServe(WEBSERVERPORT, r)
}
