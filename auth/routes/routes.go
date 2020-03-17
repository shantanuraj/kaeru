package routes

import (
	"net/http"
	v1 "sixth-io/kaeru/auth/routes/v1"

	"github.com/labstack/echo/v4"
)

var (
	v1Routes = []v1.AuthRoute{
		&v1.Signup{},
		&v1.Login{},
	}
)

// Add adds the route handler map to the echo instance
func Add(e *echo.Echo) {
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
