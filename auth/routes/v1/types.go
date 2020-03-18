package v1

import (
	"github.com/labstack/echo/v4"
)

// AuthRoute defines struct for handling routes
type AuthRoute interface {
	// Specifies if the URL should be validated with the middleware or not
	Private() bool
	// Method specifies the HTTP method to listen on
	Method() string
	// URL to bind the handler to
	URL() string
	// Handler to process the request for the URL
	Handler() echo.HandlerFunc
}
