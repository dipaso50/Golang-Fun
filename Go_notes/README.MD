
# Simple aplicación para gestionar notas de texto
----------------------------------------
La aplicación constará de un backend que ofrece servicios rest desarrollado con Golang y el front desarrollado con Angular.

Esta aplicación está basada en el ejemplo mostrado en la siguiente web:
[Developing golang and angular apps](https://auth0.com/blog/developing-golang-and-angular-apps-part-1-backend-api/)

# BackEnd ([Golang](https://golang.org/) + [Gorilla Mux](https://github.com/gorilla/mux))
----------------------------------------
Por simplicidad se ha decidido guardar las notas en memoria en un map de tipo:

```go
map[string]Note // el key string del map es un identificador único de cada objeto de tipo Note

type Note struct {
	ID      string `json:"id"`
	Msg     string `json:"message"`
	Ranking int    `json:"ranking"`
}
```

También se ha decidido definir una interfaz para permitir "inyectar" a los handler el objeto que terminará persistiendo las notas, aunque en este caso se ha decidido hacerlo en memoria, nada impide que más adelante se pueda guardar en base de datos o en fichero esa información.

```go
//NoteWarehouse operations that all warehouse must implement
type NoteWarehouse interface {
	Add(note Note) bool
	Delete(noteID string) bool
	GetAll() []Note
	GetById(noteID string) Note
}
```

Luego ya solo es cuestión de pasar la instancia correspondiente a los handler:

```go
func main() {
	r := mux.NewRouter()

	currentWarehouse = memorywarehouse.NewMemoryWarehouse() // save the information in memory

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

```

Ejecutar el backend:
```sh
$ cd $GOPATH
$ go run Go_notes/server/main.go  
``` 
 

# Ejemplos
```sh
//Get all notes
$ curl http://localhost:8080/getAll
[{"id":"bkh57ms4spfr3m9hr7ng","message":"Este es un test","ranking":0},{"id":"bkh57ms4spfr3m9hr7o0","message":"Esto es otro test","ranking":0}]

//Get one note with id = bkh57ms4spfr3m9hr7o0 
$ curl http://localhost:8080/getById/bkh57ms4spfr3m9hr7o0
{"id":"bkh57ms4spfr3m9hr7o0","message":"Esto es otro test","ranking":0}

//Add a note
$ curl -s -XPOST -d'{"id":"","message":"este es otro message from curl","ranking":1}' http://localhost:8080/add

//Delete a note with id = bkh57ms4spfr3m9hr7ng
$ curl -i -X DELETE http://localhost:8080/delete/bkh57ms4spfr3m9hr7ng

```
 

# FrontEnd ([Angular](https://angular.io/))
Pendiente implementar.

# Licencia
----------------------------------------
MIT 
