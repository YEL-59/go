package models

import "time"

// Task represents a to-do list item
type Task struct {
    ID      int       `json:"id" gorm:"primary_key;auto_increment"`
    Title   string    `json:"title"`
    DueDate time.Time `json:"due_date"`
    Status  string    `json:"status"`
}
