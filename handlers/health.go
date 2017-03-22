package handlers

import (
	"net/http"

	"github.com/kkirsche/echo_cronmon/models"
	"github.com/labstack/echo"
)

// GetHealth endpoint returns health information for the API
func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Health{
		Healthy: true,
		Message: "success",
	})
}
