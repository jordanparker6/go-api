package main

import (
	"github.com/jordanparker6/go-api/src/api"
	"github.com/jordanparker6/go-api/src/core"
	"github.com/jordanparker6/go-api/src/models"
	"github.com/jordanparker6/go-api/src/services"
)

// * API Start Up Loader

// load
// builds and returns a dependency container
// loaded with any singleton instances and executes
// any pre server startup logic.
func load() *services.Container {
	container := services.NewContainer()
	container.DB.Connect()
	container.DB.Migrate(models.All)
	return container
}

// * OpenAPI Documentation & Main Method

// @title Golang Example API
// @version 1.0
// @description This is boilerplate Golang REST API. Please feel free to contribute to this boilerplate.

// @termsOfService https://github.com/jordanparker6/go-api/LICENSE
// @contact.name Jordan Parker
// @contact.url https://github.com/jordanparker6/go-api
// @contact.email jordanhparker6@gmail.com
// @license.name MIT Licence
// @license.url https://github.com/jordanparker6/go-api/LICENSE

// @host localhost:8000
// @BasePath /api/v1
// @query.collection.format multi
// @schemes http https

// @tag.name Core
// @tag.description Core API enpoints for documentation and health checks.
// @tag.name User
// @tag.description User API endpoints for CRUD operations on user records.

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	container := load()
	api.Run(
		core.Config.Server.Host,
		core.Config.Server.Port,
		container,
	)
}
