package routes

import (
	"net/http"
	v1 "sixth-io/kaeru/auth/routes/v1"
	"sixth-io/kaeru/db"

	"github.com/labstack/echo/v4"
)

// Add adds the route handler map to the echo instance
func Add(e *echo.Echo, db *db.Database) {

	v1Routes := []v1.AuthRoute{
		v1.NewSignup(db),
		v1.NewLogin(db),
	}

	for _, config := range v1Routes {
		addv1Route(config, e)
	}
}

func addv1Route(route v1.AuthRoute, e *echo.Echo) {
	methodBinder := e.GET
	switch route.Method() {
	case http.MethodGet:
		methodBinder = e.GET
	case http.MethodPost:
		methodBinder = e.POST
	case http.MethodPut:
		methodBinder = e.PUT
	case http.MethodPatch:
		methodBinder = e.PATCH
	case http.MethodDelete:
		methodBinder = e.DELETE
	}

	methodBinder(route.URL(), route.Handler())
}
