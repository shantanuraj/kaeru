package main

import (
	"sixth-io/kaeru/auth/middleware"
	"sixth-io/kaeru/auth/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	middleware.Add(e)
	// Routes
	routes.Add(e)

	// Start server
	e.Logger.Fatal(e.Start(":4551"))
}
