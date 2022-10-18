package main

import (
	"fmt"
	"os"

	"golang-web-service/assignment-2/lib"
	"golang-web-service/assignment-2/routes"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		os.Exit(1)
	}
}

func main() {
	lib.InitDatabase()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r := routes.CreateRouter()
	r.Run(port)
}
