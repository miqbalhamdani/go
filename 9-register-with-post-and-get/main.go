package main

import (
	service "golang-web-service/9-register-with-post-and-get/services"
	"net/http"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	http.HandleFunc("/register", userSvc.RegisterHandler)
	http.HandleFunc("/user", userSvc.GetUserHandler)
	http.ListenAndServe("localhost:8080", nil)
}
