package main

import (
	"golang-chat/initDB"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	initDB.InitDb()

	

	router.GET("/api/hello", func(c *gin.Context) {
		c.IndentedJSON(http.StatusAccepted, gin.H{"data": "Hello World!", "message": "Success"})
	})

	router.Run("localhost:8080")
}
