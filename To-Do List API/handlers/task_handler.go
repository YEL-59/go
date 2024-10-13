package handlers

import (
    "net/http"
    "todo-api/db"
    "todo-api/models"
    "github.com/gin-gonic/gin"
)

// GetTasks fetches all tasks from the database
func GetTasks(c *gin.Context) {
    var tasks []models.Task
    db.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}

// CreateTask creates a new task in the database
func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&task)
    c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task
func UpdateTask(c *gin.Context) {
    var task models.Task
    id := c.Param("id")
    if err := db.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&task)
    c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
    var task models.Task
    id := c.Param("id")
    if err := db.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    db.DB.Delete(&task)
    c.JSON(http.StatusNoContent, nil)
}
