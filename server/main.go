package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var tasksCollection *mongo.Collection

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

func handleAdd(c *gin.Context) {
	// Read the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error, e.g., log it or return an error response
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	// Convert the request body to a string and print it
	requestBody := string(body)
	fmt.Println("Adding data:", requestBody)

	// You can now use the requestBody as needed
	c.String(http.StatusOK, "Adding data: "+requestBody)
}

func handleEdit(c *gin.Context) {
	// Read the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error, e.g., log it or return an error response
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	// Convert the request body to a string and print it
	requestBody := string(body)
	fmt.Println("Editing data:", requestBody)

	// You can now use the requestBody as needed
	c.String(http.StatusOK, "Editing data: "+requestBody)
}

func handleDelete(c *gin.Context) {
	data := c.PostForm("data")
	fmt.Println("Deleting data:", data)
	c.String(http.StatusOK, "Deleting data: "+data)
}

func handleGetTasks(c *gin.Context) {
	fmt.Println("Retrieving tasks...")
	c.String(http.StatusOK, "Retrieving tasks...")
}

func defineRoutes(r *gin.Engine) {
	r.Use(corsMiddleware())

	r.POST("/add", handleAdd)
	r.POST("/edit", handleEdit)
	r.POST("/delete", handleDelete)
	r.GET("/tasks", handleGetTasks)
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	tasksCollection = client.Database("TaskManagerDB").Collection("Tasks")

	r := gin.Default()

	defineRoutes(r)

	r.Run(":8081")
}
