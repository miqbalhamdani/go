package controllers

import (
	"fmt"
	"golang-web-service/assignment-3/helpers"
	"golang-web-service/assignment-3/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	weather := models.Weather{
		Water: helpers.GenerateRandomNumberRange(),
		Wind:  helpers.GenerateRandomNumberRange(),
	}

	status := helpers.GetStatus(weather)

	dataJson := models.Data{
		Status: weather,
	}
	fmt.Println(helpers.Stringify(dataJson))

	// Create Folder
	folder := "storage/"
	filename := "weather.json"
	fullPath := folder + filename
	helpers.CreateFolder(folder)

	// Write File
	helpers.WriteFile(dataJson, fullPath)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"status": status,
	})

	return
}
