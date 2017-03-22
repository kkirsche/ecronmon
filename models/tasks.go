package models

import (
	"time"

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
func GetTask(id int, db *gorm.DB) Task {
	var task Task
	db.Find(&task, id)
	return task
}

// CreateTask is used to create a new task object in the database
func CreateTask(task Task, db *gorm.DB) Task {
	db.Save(&task)
	return task
}

// DeleteTask is used to delete a task object in the database
func DeleteTask(id int, db *gorm.DB) {
	var task Task
	db.Find(&task, id)
	db.Delete(&task)
}
