package service

import (
	"sesi_11/entity"
	"sesi_11/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepo}

func TestProductServiceGetOneProductByIDNotFound(t *testing.T) {
	productRepo.Mock.On("FindById", 1).Return(nil)

	product, err := productService.GetOneProductById(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "Error message should be 'product not found'")
}

func TestProductServiceGetOneProductByID(t *testing.T) {
	productData := entity.Product{
		ID:   2,
		Name: "Product 1",
	}

	productRepo.Mock.On("FindById", 2).Return(productData)

	product, err := productService.GetOneProductById(2)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productData.ID, product.ID, "Product ID should be 1")
	assert.Equal(t, productData.Name, product.Name, "Product name should be 'Product 1'")
	assert.Equal(t, &productData, product, "Product should be equal")
}
