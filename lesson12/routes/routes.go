package routes

import (
	"example.com/rest-API/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticate := server.Group("/")
	authenticate.POST("/events", middlewares.Authenticate, createEvents)

	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/signup", signup)
	authenticate.POST("/login", login)
}
