package controller

import (
	"fmt"
	"sesi_6/pkg/middleware"
	"sesi_6/pkg/models"
	"sesi_6/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Routes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")

	authGroup.POST("/register", c.Register)
	authGroup.POST("/login", c.Login)
	authGroup.GET("/me", middleware.VerifyToken, c.Me)
}

func (c *UserController) Login(ctx *gin.Context) {
	var request models.LoginRequest

	err := ctx.ShouldBind(&request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Login(request)

	if err != nil {
		ctx.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}

func (c *UserController) Register(ctx *gin.Context) {
	var request models.RegisterRequest

	err := ctx.ShouldBind(&request)

	fmt.Printf("%+v", request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Register(request)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}

func (c *UserController) Me(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	userData := user.(jwt.MapClaims)

	fmt.Printf("username %v", userData["username"])

	ctx.JSON(200, gin.H{
		"data": "berhasil",
	})
}
