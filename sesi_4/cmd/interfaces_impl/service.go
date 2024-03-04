package main

type ProductService struct {
	ProductRepository ProductRepository
}

func (s ProductService) FindAll() []Product {
	return s.ProductRepository.FindAll()
}

func (s ProductService) FindById(id int) Product {
	return s.ProductRepository.FindById(id)
}

func (s ProductService) Save(product Product) Product {
	return s.ProductRepository.Save(product)
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{ProductRepository: repo}
}
