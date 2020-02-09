package producer

import (
	"Go_rabbitMQ_test/Go_rabbit_producer/customtypes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

//Add a message to the message Broker
func Add() func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		var msg customtypes.Message

		json.NewDecoder(r.Body).Decode(&msg)

		var err error

		//Generamos un único ID para cada mensaje
		msg.ID, err = generateUniqueID()

		failOnError(err, "Error generando el ID")

		fmt.Printf("Producer. Sending message to RabbitMQ :(%s) \n", msg)

		produceMessage(msg)
	}

	return f
}

func produceMessage(msg customtypes.Message) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Producer: Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	bodyMsg, err := json.Marshal(msg)

	failOnError(err, "Failed encoding")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(bodyMsg),
		})

	failOnError(err, "Failed to publish a message")
}

func generateUniqueID() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
