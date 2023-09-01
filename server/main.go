package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	// Parse the request body into a map or struct (based on your data structure)
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.String(http.StatusBadRequest, "Error parsing request body")
		return
	}

	// Generate a unique ID for the new document
	id := primitive.NewObjectID()
	requestData["_id"] = id

	// Insert the data into the MongoDB collection
	insertResult, insertErr := tasksCollection.InsertOne(context.TODO(), requestData)
	if insertErr != nil {
		c.String(http.StatusInternalServerError, "Error inserting data into MongoDB")
		return
	}

	// Retrieve the ID of the newly inserted document
	insertedID := insertResult.InsertedID

	// You can now use the insertedID and the requestData as needed
	c.String(http.StatusOK, fmt.Sprintf("Data added successfully with ID: %v", insertedID))
}

func handleEdit(c *gin.Context) {
	// Read the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	// Parse the request body into a map or struct (based on your data structure)
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.String(http.StatusBadRequest, "Error parsing request body")
		return
	}

	// Ensure the request includes a valid _id field
	idStr, ok := requestData["_id"].(string)
	if !ok {
		c.String(http.StatusBadRequest, "Invalid request data, missing '_id'")
		return
	}

	// Convert the _id string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid _id format")
		return
	}

	// Remove the _id field from the requestData as we cannot update it
	delete(requestData, "_id")

	// Create the update query
	update := bson.D{
		{"$set",
			requestData,
		},
	}

	// Execute the UpdateByID() function with the filter and update query
	result, err := tasksCollection.UpdateByID(context.TODO(), objectID, update)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating data")
		return
	}

	// Handle the update success response
	c.String(http.StatusOK, fmt.Sprintf("Editing data: Updated %d document(s)", result.ModifiedCount))
}

func handleDelete(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.String(http.StatusBadRequest, "Error parsing request body")
		return
	}

	id, ok := requestData["_id"].(string)
	if !ok {
		c.String(http.StatusBadRequest, "Invalid request data, missing '_id'")
		return
	}

	// Convert the "_id" string to an ObjectID (assuming MongoDB uses ObjectID as primary keys)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid '_id' format")
		return
	}

	// Create a filter to find the document with the specified "_id"
	filter := bson.M{"_id": objectID}

	// Delete the item from the MongoDB collection based on the filter
	deleteResult, err := tasksCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting data from MongoDB")
		return
	}

	// Check if any documents were deleted
	if deleteResult.DeletedCount == 0 {
		c.String(http.StatusNotFound, "No matching document found for deletion")
		return
	}

	c.String(http.StatusOK, "Deleted data with ID: "+id)
}

func handleGetTasks(c *gin.Context) {
	// Retrieve all tasks from the MongoDB collection
	cursor, err := tasksCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving tasks from MongoDB")
		return
	}
	defer cursor.Close(context.TODO())

	// Create a slice to store the retrieved tasks
	var tasks []bson.M

	// Iterate through the cursor and decode the documents into tasks
	for cursor.Next(context.TODO()) {
		var task bson.M
		if err := cursor.Decode(&task); err != nil {
			c.String(http.StatusInternalServerError, "Error decoding task document")
			return
		}
		tasks = append(tasks, task)
	}

	// Print the retrieved tasks (you can customize this output as needed)
	for _, task := range tasks {
		fmt.Println("Retrieved task:", task)
	}

	// You can now use the retrieved tasks as needed
	c.JSON(http.StatusOK, tasks)
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
