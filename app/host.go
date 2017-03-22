package app

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// Host represents a network host
type host struct {
	ID        int        `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	IP        string     `json:"ip_address"`
}

var (
	hosts = map[int]*host{}
	seq   = 1
)

func migrateHost(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS hosts(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			ip_address VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		return err
	}

	return nil
}

// CreateHost is used to create a new host object
func CreateHost(c echo.Context) error {
	h := new(host)
	if err := c.Bind(h); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ErrorResponse{
		Message: "ID was set. This is not an allowed parameter",
	})
}

// GetHosts is used to retrieve a list of hosts
func GetHosts(c echo.Context) error {
	return c.JSON(http.StatusOK, hosts)
}

// GetHost is used to get a single host
func GetHost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "ID was not a number",
		})
	}
	if hosts[id] != nil {
		return c.JSON(http.StatusOK, hosts[id])
	}
	return c.JSON(http.StatusOK, ErrorResponse{
		Message: "No results found",
	})
}

// UpdateHost is used to update the values of a single host
func UpdateHost(c echo.Context) error {
	h := new(host)
	if err := c.Bind(h); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	hosts[id].Name = h.Name
	hosts[id].UpdatedAt = time.Now()
	hosts[id].IP = h.IP
	return c.JSON(http.StatusOK, hosts[id])
}

// DeleteHost is used to delete a single host
func DeleteHost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(hosts, id)
	return c.NoContent(http.StatusNoContent)
}
