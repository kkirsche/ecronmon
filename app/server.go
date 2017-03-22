package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run is used to setup and execute the Echo server
func Run() {
	db := initDB("storage.db")
	if db == nil {
		return
	}
	err := migrate(db)
	if err != nil {
		return
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	g := e.Group("/api")

	// Routes
	g.GET("/hosts", GetHosts)
	g.GET("/hosts/:id", GetHost)
	g.POST("/hosts", CreateHost)
	g.PUT("/hosts/:id", UpdateHost)
	g.DELETE("/hosts/:id", DeleteHost)

	e.Logger.Fatal(e.Start(":8080"))
}
