package webserver

import (
	"fmt"
	"golang-web-service/10-web-server/webserver/controllers"
	"log"
	"net/http"
)

const PORT = ":4000"

func Start() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/users", controllers.GetUsersHandler)
	http.HandleFunc("/users/create", controllers.CreateUserHandler)
	log.Println("Server running at port", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}
