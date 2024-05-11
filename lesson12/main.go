package main

import (
	"example.com/rest-API/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Run(":8080")
}
