package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func defineRoutes(r *gin.Engine) {

	r.Use(corsMiddleware())

	r.POST("/add", func(c *gin.Context) {
		data := c.PostForm("data")
		fmt.Println("Here")
		c.String(http.StatusOK, "Adding data: "+data)
	})

	r.POST("/edit", func(c *gin.Context) {
		data := c.PostForm("data")
		c.String(http.StatusOK, "Editing data: "+data)
	})

	r.POST("/delete", func(c *gin.Context) {
		data := c.PostForm("data")
		c.String(http.StatusOK, "Deleting data: "+data)
	})

	r.GET("/tasks", func(c *gin.Context) {
		c.String(http.StatusOK, "Retrieving tasks...")
	})
}

func main() {
	r := gin.Default()

	defineRoutes(r)

	r.Run(":8081")
}
