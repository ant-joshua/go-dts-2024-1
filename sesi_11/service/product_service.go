package service

import (
	"errors"
	"sesi_11/entity"
	"sesi_11/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProductById(id int) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
