package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {

}
