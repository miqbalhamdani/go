package main

import (
	"golang-web-service/assignment-3/routes"
	"log"
)

const PORT = ":3000"

func main() {
	router := routes.WeatherHttpHandler()
	log.Fatal(router.Run(PORT))
}
