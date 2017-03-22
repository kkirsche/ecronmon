package app

import (
	"github.com/kkirsche/echo_cronmon/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunAPI is used to setup and execute the Echo API server
func RunAPI() {
	db := initDB("storage.db")
	if db == nil {
		return
	}
	migrate(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// API routes
	g := e.Group("/api")
	g.GET("/tasks", handlers.IndexTask(db))
	g.GET("/tasks/:id", handlers.ShowTask(db))
	g.POST("/tasks", handlers.CreateTask(db))
	// g.PUT("/tasks/:id", handlers.UpdateTask(db))
	g.DELETE("/tasks/:id", handlers.DestroyTask(db))

	e.Logger.Fatal(e.Start(":8080"))
}

// RunFrontend starts the Vue.js front end application
func RunFrontend() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// // Non-API Routes
	e.File("/", "public/index.html")

	e.Logger.Fatal(e.Start(":8080"))
}
