package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Convert the notFoundResponse() helper to a http.Handler using the
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//Below is an example to help properly form your routes(You can remove when you are done)
	// router.HandlerFunc(http.MethodGet, "/v1/animals", app.listAnimalsHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/animals", app.createAnimalHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/animals/:id", app.showAnimalHandler)
	// router.HandlerFunc(http.MethodPatch, "/v1/animals/:id", app.updateAnimalHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/animals/:id", app.deleteAnimalHandler)

	// Return the httprouter instance.
	return router
}
