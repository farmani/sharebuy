package main

import (
	echoPrometheus "github.com/labstack/echo-contrib/echoprometheus"

	"github.com/labstack/echo/v4"
)

// routes is our main application's router.
func (app *application) routes(e *echo.Echo) {
	app.bundleMiddleware(e)

	//e.Router().NotFound = http.HandlerFunc(app.notFoundResponse)

	// Convert app.methodNotAllowedResponse helper to a http.Handler and set it as the custom
	// error handler for 405 Method Not Allowed responses
	//router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// healthcheck
	e.GET("/v1/health", app.healthHandler)

	// adds middleware to gather metrics from those endpoints and expose them on /metrics
	e.GET("/metrics", echoPrometheus.NewHandler())

	// Users handlers
	e.POST("/v1/users", app.registerStartHandler)
	e.PUT("/v1/users/activated", app.registerFinishHandler)
	// e.POST("/v1/users/login", app.createAuthenticationTokenHandler)
	// e.POST("/v1/users/logout", app.createAuthenticationTokenHandler)
	// e.POST("/v1/users/forget", app.createAuthenticationTokenHandler)

	// Movies handlers. Note, that these movie endpoints use the `requireActivatedUser` middleware.
	// e.GET("/v1/products", app.requirePermissions("movies:read", app.listMoviesHandler))
	// e.POST("/v1/products", app.requirePermissions("movies:write", app.createMovieHandler))
	// e.GET("/v1/products/:id", app.requirePermissions("movies:read", app.showMovieHandler))
	// e.PATCH("/v1/products/:id", app.requirePermissions("movies:write", app.updateMovieHandler))
	// e.DELETE("/v1/products/:id", app.requirePermissions("movies:write", app.deleteMovieHandler))

	// Wrap the router with the panic recovery middleware and rate limit middleware.
	//return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
