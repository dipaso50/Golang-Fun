
# Microservicios que utilizan RabbitMQ para comunicarse, ejemplo basado en el tutorial de [RabbitMQ para Golang](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)
La arquitectura de la aplicación es sencilla, se define un Web Services REST(Producer) que aceptará peticiones y las remitirá al message broker para que los consumidores atiendan la petición y guarden en mensaje en una base de datos MySql.

![Alt text](appSchema.png?raw=true "App Schema")


## Ejecución
--------
1. Levantar los servicios de RabbitMQ y MySql con Docker.
2. Lanzar el Productor y el número deseado de consumidores.
- Productor 
```bash
$ go run $GOPATH/src/Go_rabbitMQ_test/Go_rabbit_producer/main.go
```
- Consumidor, se pueden lanzar cuantos hagan falta
```bash
$ go run $GOPATH/src/Go_rabbitMQ_test/Go_rabbit_consumer/main.go
```


Teniendo ya todo levantado un funcionando, solo es cuestión de hacer una petición de Post al productor con cualquier cliente, ya sea Postman, Curl, etc.
Por ejemplo, con el siguiente script de Python podemos llamar a nuestro Producer N veces.
```python
import requests
import json, sys

repeticiones = int(sys.argv[1])
count = 0
while (count < repeticiones):    
    msg = "Message " + str(count)
    count = count + 1
    r = requests.post("http://localhost:3000/", json.dumps({'Data': msg}))  

```

Para simular 2000 llamadas a nuestro Producer:
```bash
$ python3  $GOPATH/src/Go_rabbitMQ_test/CallProducer.py 2000 
```

## Mysql y RabbitMQ
--------
Por simplificar se opta ejecutar el servidor de Mysql y RabbitMQ con Docker.

[MySql con Docker](https://hub.docker.com/_/mysql/) 
```s
$ docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
```

Nos conectamos al contenedor Mysql, creamos la base de datos **rabbitmq_message** y la tabla **Message**.
```bash
root@07690317b2f6:/# mysql -p     
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10
Server version: 8.0.19 MySQL Community Server - GPL

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show tables
    -> ;
ERROR 1046 (3D000): No database selected
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.01 sec)

mysql> create database rabbitmq_message;
Query OK, 1 row affected (0.02 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| rabbitmq_message   |
| sys                |
+--------------------+
5 rows in set (0.01 sec)

mysql> CREATE TABLE IF NOT EXISTS Message \
    -> (ID VARCHAR(255) PRIMARY KEY,\
    -> Msg VARCHAR(255) NOT NULL, \
    -> created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
ERROR 1046 (3D000): No database selected
mysql> use rabbitmq_message;
Database changed
mysql> CREATE TABLE IF NOT EXISTS Message \
    -> (ID VARCHAR(255) PRIMARY KEY,\
    -> Msg VARCHAR(255) NOT NULL, \
    -> created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
Query OK, 0 rows affected (0.05 sec)

mysql> show tables;
+----------------------------+
| Tables_in_rabbitmq_message |
+----------------------------+
| Message                    |
+----------------------------+
1 row in set (0.00 sec)

mysql> 
```

Es necesario preguntar a Docker la ip que usaremos para conectarnos a MySql. Esta ip es la que usaran los consumidores para conectarse.
```bash
diego@vant-N14xWU:~$ docker inspect some-mysql | grep IPAddress
            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",

```

**RabbitMQ con Docker**
```s
$ docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

# Producer
--------
El productor realiza tres acciones principalmente:

* Escucha peticiones de clientes. 
```go
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", TestAlive).Methods("GET")
	r.HandleFunc("/", producer.Add()).Methods("POST")

	log.Fatal(http.ListenAndServe(WEBSERVERPORT, r))
}
``` 
* Decodifica el json de la petición y genera un ID para cada mensaje.

```go
func Add() func(w http.ResponseWriter, r *http.Request) {

	f := func(w http.ResponseWriter, r *http.Request) {
		var msg customtypes.Message

		json.NewDecoder(r.Body).Decode(&msg)

		var err error

		//Generamos un único ID para cada mensaje
		msg.ID, err = generateUniqueID()

		failOnError(err, "Error generando el ID")

		fmt.Printf("Produciendo Mensaje :(%s) \n", msg)

		produceMessage(msg)
	}

	return f
}
```
* Conecta con RabbitMQ para enviar la petición.
```go
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
```


# Consumer
--------
El consumidor se conecta con RabbitMQ para escuchar los mensajes y guardar la información en MySql.
```go
//ConsumeForever escucha continuamente del message broker
func ConsumeForever() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Consumer: Failed to connect to RabbitMQ")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var msg customtypes.Message

			//json.NewDecoder(d.Body).Decode(&msg)
			err := json.Unmarshal([]byte(d.Body), &msg)

			failOnError(err, "Error decoding")

			log.Printf("Received a message: %s", d.Body)

			saveindatabase(msg)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func saveindatabase(message customtypes.Message) {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(172.17.0.3:3306)/rabbitmq_message")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO Message VALUES ( '" + message.ID + "', '" + message.Data + "' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
```
 

# Licencia
MIT 
