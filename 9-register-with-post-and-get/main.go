package main

import (
	"net/http"
	service "session-1/9-register-with-post-and-get/services"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	http.HandleFunc("/register", userSvc.RegisterHandler)
	http.HandleFunc("/user", userSvc.GetUserHandler)
	http.ListenAndServe("localhost:8080", nil)
}
