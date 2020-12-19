// The API's controller methods
// perform the API logic and build the
// required services using the DI container
// OpenAPI documentation for the endpoints is
// provided against each controller method.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanparker6/go-api/src/services"
)

type controller struct {
	depends *services.Container
}

// NewController
// builds a controller struct with the global DI container.
func NewController(depends *services.Container) controller {
	controller := new(controller)
	controller.depends = depends
	return *controller
}

// * Core Controller Methods

// @Summary Test API
// @Description An endpoint to test the API's connectivity.
// @ID test-api
// @Tags Core
// @Router /test [get]
// @Produce  json
// @Success 200 {object} core.HTTPSuccess
// @Failure 404 {object} core.HTTPError
// @Failure 500 {object} core.HTTPError
func (c *controller) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

// @Summary OpenAPI
// @Description Fetches the OpenAPI json specification for the API.
// @ID open-api
// @Tags Core
// @Router /openapi.json [get]
// @Produce  json
// @Success 200
// @Failure 404 {object} core.HTTPError
// @Failure 500 {object} core.HTTPError
func (c *controller) GetOpenAPI(ctx *gin.Context) {
	ctx.File("docs/swagger.json")
}

// GetReDocs
// Renders the ReDoc HTML for docs. OpenAPI spec not required.
func (c *controller) GetReDocs(ctx *gin.Context) {
	html := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>ReDoc</title>
			<!-- needed for adaptive design -->
			<meta charset="utf-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

			<!--
			ReDoc doesn't change outer page styles
			-->
			<style>
			body {
				margin: 0;
				padding: 0;
			}
			</style>
		</head>
		<body>
			<redoc spec-url='/api/v1/openapi.json'></redoc>
			<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
		</body>
	</html>
	`
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(html))
}
