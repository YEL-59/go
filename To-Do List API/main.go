package main

import (
    "todo-api/db"
    "todo-api/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    db.InitializeDatabase()

    // Set up router
    r := gin.Default()

    // Define routes
    r.GET("/tasks", handlers.GetTasks)
    r.POST("/tasks", handlers.CreateTask)
    r.PUT("/tasks/:id", handlers.UpdateTask)
    r.DELETE("/tasks/:id", handlers.DeleteTask)

    // Start the server
    r.Run(":8080")
}
