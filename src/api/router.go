package api

import (
	"fmt"

	"github.com/jordanparker6/go-api/src/api/controller"
	"github.com/jordanparker6/go-api/src/services"

	"github.com/gin-gonic/gin"
)

// Run
// builds the router, attaches enpoints and runs the server.
func Run(host string, port string, depends *services.Container) {
	router := gin.Default()
	api := router.Group("/api/v1")
	c := controller.NewController(depends)
	{
		users := api.Group("/user")
		{
			users.GET(":id", c.GetUser)
			users.POST("create", c.CreateUser)
			users.DELETE(":id", c.DeleteUser)
		}
		api.GET("/docs", c.GetReDocs)
		api.GET("/openapi.json", c.GetOpenAPI)
		api.GET("/test", c.Test)
	}
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
