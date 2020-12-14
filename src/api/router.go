package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()
var api *gin.RouterGroup = router.Group("v1")

func Build() {
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})
}

func Run(host string, port string) {
	Build()
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
