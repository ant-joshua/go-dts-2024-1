package main

import (
	"encoding/json"
	"net/http"
)

type ProductController struct {
	ProductService ProductService
}

func (c ProductController) Get(res http.ResponseWriter, req *http.Request) {
	products := c.ProductService.FindAll()
	json.NewEncoder(res).Encode(products)
}

func NewProductController(service ProductService) ProductController {
	return ProductController{ProductService: service}
}

func main() {

	productRepo := ProductRepositoryMongo{}
	productService := NewProductService(productRepo)
	productController := NewProductController(productService)

	http.HandleFunc("/products", productController.Get)

}
