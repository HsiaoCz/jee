package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/sayhello", SayHello)
	r.Run("127.0.0.1:4002")
}

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello",
	})
}
