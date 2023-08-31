package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function to define route handlers
func defineRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.POST("/add", func(c *gin.Context) {
		data := c.PostForm("data")
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
