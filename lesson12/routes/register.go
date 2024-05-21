package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
}

func cancelRegistration() {

}
