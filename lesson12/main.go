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
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}