package main

import (
	"log"

	"golang-web-service/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Mygram API
// @version 1.0
// @description MyGram API for Final Project Scalable Web Service with Golang by Hacktiv8

// @BasePath /

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
