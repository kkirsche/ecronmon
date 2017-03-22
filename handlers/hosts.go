package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kkirsche/echo_cronmon/models"
	"github.com/labstack/echo"
)

// IndexHost endpoint returns a list of all hosts
func IndexHost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch hosts using our model
		hosts := models.GetHosts(db)

		return c.JSON(http.StatusOK, hosts)
	}
}

// ShowHost endpoint returns a single host
func ShowHost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch individual host using our model
		id, _ := strconv.Atoi(c.Param("id"))
		host := models.GetHost(db, id)
		bhost := models.Host{}
		if host == bhost {
			return c.JSON(http.StatusOK, models.ErrorResponse{
				Status:  http.StatusNotFound,
				Title:   "Host not found",
				Message: fmt.Sprintf("Host with id %d could not be found", id),
			})
		}
		return c.JSON(http.StatusOK, host)
	}
}

// CreateHost endpoint
func CreateHost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new host
		var host models.Host
		// Map incoming JSON body to the new Host
		c.Bind(&host)
		// Add a host using our new model
		host = models.CreateHost(db, host)

		return c.JSON(http.StatusCreated, host)
	}
}

// UpdateHost endpoint
func UpdateHost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		host := models.GetHost(db, id)
		c.Bind(&host)
		models.UpdateHost(db, host)

		return c.JSON(http.StatusCreated, host)
	}
}

// DestroyHost endpoint
func DestroyHost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		host := models.GetHost(db, id)
		models.DeleteHost(db, host)

		return c.JSON(http.StatusOK, host)
	}
}
