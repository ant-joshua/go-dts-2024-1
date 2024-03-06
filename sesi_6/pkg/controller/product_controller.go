package controller

import (
	"fmt"
	"sesi_6/pkg/models"
	"sesi_6/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (p *ProductController) Routes(r *gin.RouterGroup) {
	routeGroup := r.Group("/products")

	routeGroup.GET("", p.GetAllProduct)
	routeGroup.POST("", p.CreateProduct)
	routeGroup.GET("/:id", p.GetProductByID)
	routeGroup.PUT("/:id", p.UpdateProduct)
	routeGroup.DELETE("/:id", p.DeleteProduct)
}

func (p *ProductController) GetAllProduct(ctx *gin.Context) {

	var request models.GetListProductRequest

	err := ctx.ShouldBindQuery(&request)

	fmt.Printf("%+v", request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	products, err := p.service.GetAllProduct(request)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var request models.CreateProductRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.CreateProduct(request)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, product)
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	convertID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.GetProductByID(convertID)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, product)
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	convertID, err := strconv.Atoi(id)

	fmt.Println(convertID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var request models.UpdateProductRequest

	err = ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.UpdateProduct(convertID, request)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(200, product)
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	convertID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = p.service.DeleteProduct(convertID)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product deleted"})
}
