package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kkirsche/echo_cronmon/models"
	"github.com/labstack/echo"
)

// IndexTask endpoint returns a list of all tasks
func IndexTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch tasks using our model
		tasks := models.GetTasks(db)

		return c.JSON(http.StatusOK, tasks)
	}
}

// ShowTask endpoint returns a single task
func ShowTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch individual task using our model
		id, _ := strconv.Atoi(c.Param("id"))
		task := models.GetTask(id, db)
		btask := models.Task{}
		if task == btask {
			return c.JSON(http.StatusOK, models.ErrorResponse{
				Status:  http.StatusNotFound,
				Title:   "Task not found",
				Message: fmt.Sprintf("Task with id %d could not be found", id),
			})
		}
		return c.JSON(http.StatusOK, task)
	}
}

// CreateTask endpoint
func CreateTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map incoming JSON body to the new Task
		c.Bind(&task)
		// Add a task using our new model
		task = models.CreateTask(task, db)

		return c.JSON(http.StatusCreated, task)
	}
}

// UpdateTask endpoint
// func UpdateTask(db *gorm.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		task := models.GetTask(id, db)
// 		c.Bind(&task)
// 		task, err := models.UpdateTask(task, db)
//
// 		return c.JSON(http.StatusCreated, models.Task{ID: id})
// 	}
// }

// DestroyTask endpoint
func DestroyTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		models.DeleteTask(id, db)

		return c.JSON(http.StatusOK, models.Deleted{ID: id})
	}
}
