package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-API/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event."})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user."})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered."})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventID

	event.cancelRegistration(userID)
}
