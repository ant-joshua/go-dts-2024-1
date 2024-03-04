package routers

import (
	"sesi_6/pkg/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// API
	api := r.Group("/api")

	employeeController := controller.EmployeeController{}
	employeeController.Routes(api)

	return r
}
