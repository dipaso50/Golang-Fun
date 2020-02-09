package main

import (
	"Go_rabbitMQ_test/Go_rabbit_producer/producer"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	//WEBSERVERPORT port where the application listen
	WEBSERVERPORT = ":3000"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", TestAlive).Methods("GET")
	r.HandleFunc("/", producer.Add()).Methods("POST")
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	log.Fatal(http.ListenAndServe(WEBSERVERPORT, r))
}

//TestAlive say hello word
func TestAlive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
