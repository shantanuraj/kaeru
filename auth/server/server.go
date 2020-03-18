package server

import (
	"sixth-io/kaeru/auth/middlewares"
	"sixth-io/kaeru/auth/routes"
	"sixth-io/kaeru/db"

	"github.com/labstack/echo/v4"
)

// Start boots the server instance
func Start(port string, dao *db.Database) {
	// Echo instance
	e := echo.New()

	// Middleware
	middlewares.Add(e)
	// Routes
	routes.Add(e, dao)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
