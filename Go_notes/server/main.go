package main

import (
	"Go_notes/server/handler"
	"Go_notes/server/memorywarehouse"
	"Go_notes/server/notes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var currentWarehouse notes.NoteWarehouse

const (
	//WEBSERVERPORT port where the application listen
	WEBSERVERPORT = ":3000"
)

func main() {
	r := mux.NewRouter()

	currentWarehouse = memorywarehouse.NewMemoryWarehouse()

	n := notes.NewNote("Nota 1", "Este es el primer detalle", 0)
	currentWarehouse.Add(n)

	n2 := notes.NewNote("Nota 2", "ESte es el segundo detalle de test", 0)
	currentWarehouse.Add(n2)

	r.HandleFunc("/", handler.Hello).Methods("GET")
	r.HandleFunc("/add", handler.Add(currentWarehouse)).Methods("POST")
	r.HandleFunc("/getAll", handler.GetAll(currentWarehouse)).Methods("GET")
	r.HandleFunc("/getById/{noteId}", handler.GetByID(currentWarehouse)).Methods("GET")
	r.HandleFunc("/delete/{noteId}", handler.Delete(currentWarehouse)).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "OPTIONS", "DELETE"},
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(WEBSERVERPORT, handler))
}
