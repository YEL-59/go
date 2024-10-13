package db

import (
    "gorm.io/driver/sqlite" 
    "gorm.io/gorm"
    "todo-api/models"
    "log"
)

var DB *gorm.DB

// InitializeDatabase sets up the database connection
func InitializeDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{}) 
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Auto-migrate the Task model
    DB.AutoMigrate(&models.Task{})
}
