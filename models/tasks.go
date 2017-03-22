package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Tasks is an array of Task objects
type Tasks []Task

// Task represents a cron task
type Task struct {
	ID        uint       `json:"id"`
	HostID    uint       `json:"host_id"`
	URLID     string     `json:"url_uuid"`
	Name      string     `json:"name"`
	Frequency string     `json:"frequency"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// GetTasks is used to retrieve a list of task objects from the database
func GetTasks(db *gorm.DB) Tasks {
	var tasks Tasks
	db.Find(&tasks)
	return tasks
}

// GetTask is used to retrieve a single task object from the database
func GetTask(db *gorm.DB, id int) Task {
	var task Task
	db.Find(&task, id)
	return task
}

// CreateTask is used to create a new task object in the database
func CreateTask(db *gorm.DB, task Task) Task {
	task.URLID = uuid.New().String()
	db.Save(&task)
	return task
}

// UpdateTask is used to update an existing task object in the database
func UpdateTask(db *gorm.DB, task Task) {
	db.Save(&task)
}

// DeleteTask is used to delete a task object in the database
func DeleteTask(db *gorm.DB, task Task) {
	db.Delete(&task)
}
