package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
}

func cancelRegistration() {

}
