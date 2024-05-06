package main

import (
	"net/http"

	"example.com/rest-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvent()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {

}
