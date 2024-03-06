package routers

import (
	"database/sql"
	"sesi_6/pkg/controller"
	"sesi_6/pkg/service"

	_ "sesi_6/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartServer(db *sql.DB, gorm *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API
	api := r.Group("/api")

	employeeController := &controller.EmployeeController{
		DB: db,
	}
	employeeController.Routes(api)

	productService := service.NewProductService(gorm)
	productController := controller.NewProductController(productService)
	productController.Routes(api)

	return r
}
