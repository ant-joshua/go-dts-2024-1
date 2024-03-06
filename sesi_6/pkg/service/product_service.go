package service

import (
	"fmt"
	"sesi_6/pkg/models"

	"gorm.io/gorm"
)

type ProductService struct {
	gorm *gorm.DB
}

func NewProductService(gorm *gorm.DB) *ProductService {
	return &ProductService{gorm: gorm}
}

func (p *ProductService) GetAllProduct(request models.GetListProductRequest) ([]models.Product, error) {

	var products []models.Product

	findProducts := p.gorm

	if request.Name != nil {

		findProducts = findProducts.Where("product_name ILIKE ?", fmt.Sprintf("%%%s%%%%", *request.Name))
	}

	if request.StartPrice != nil && request.EndPrice != nil {
		findProducts = findProducts.Where("product_price BETWEEN ? AND ?", *request.StartPrice, *request.EndPrice)
	}

	err := findProducts.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) CreateProduct(request models.CreateProductRequest) (*models.Product, error) {

	product := models.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	err := p.gorm.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductService) GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	err := p.gorm.Where("product_id = ?", id).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductService) UpdateProduct(id int, request models.UpdateProductRequest) (*models.Product, error) {
	product, err := p.GetProductByID(id)

	fmt.Printf("product: %+v\n", product)

	if err != nil {
		return nil, err
	}

	product.Name = request.Name
	product.Price = request.Price

	err = p.gorm.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) DeleteProduct(id int) error {
	err := p.gorm.Delete(&models.Product{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
