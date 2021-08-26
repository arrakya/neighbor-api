package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var context []Context

func main() {
	r := gin.Default()
	fmt.Println("Web API Start")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.JSON(200, gin.H{
			"docs": context,
		})
	})

	r.POST("/", func(c *gin.Context) {
		c.BindJSON(&context)

		c.JSON(200, gin.H{
			"count": len(context),
		})
	})

	r.Run()
}

// Context for file and directo information
type Context struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Contents []Context `json:"contents"`
}
