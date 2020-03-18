package routes

import (
	"sixth-io/kaeru/auth/middlewares"
	v1 "sixth-io/kaeru/auth/routes/v1"
	"sixth-io/kaeru/db"

	"github.com/labstack/echo/v4"
)

// Add adds the route handler map to the echo instance
func Add(e *echo.Echo, db *db.Database) {

	v1Group := e.Group("/v1")
	v1Routes := []v1.AuthRoute{
		v1.NewSignup(db),
		v1.NewLogin(db),
		v1.NewValidate(db),
	}

	for _, config := range v1Routes {
		addv1Route(config, v1Group, db)
	}
}

func addv1Route(route v1.AuthRoute, g *echo.Group, db *db.Database) {
	if route.Private() {
		g.Add(
			route.Method(),
			route.URL(),
			route.Handler(),
			middlewares.VerifyToken(db),
		)
	} else {
		g.Add(route.Method(), route.URL(), route.Handler())
	}
}
