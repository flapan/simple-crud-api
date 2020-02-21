package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

var users = map[int]User{}

// User sfjsojfosjfosjf
type User struct {
	Name string
	Age  int
}

// Home does something
func Home(w http.ResponseWriter, r *http.Request) {
	respond(w, "hello api")
}

// Users does something else
func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		respond(w, "get users")
	case http.MethodPost:
		user := User{
			Name: "mikkel",
			Age:  46,
		}
		users[len(users)] = user
		respond(w, user)
	}

}

func respond(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("ERR: could marshal data", err.Error())
	}
	bytesCount, err := w.Write(bytes)
	if err != nil {
		log.Println("ERR: could not write")
	}
	log.Println("Number of bytes", bytesCount)
}
