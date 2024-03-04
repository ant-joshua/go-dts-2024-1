package main

type ProductRepository interface {
	FindAll() []Product
	FindById(id int) Product
	Save(product Product) Product
}

type ProductRepositoryPostgres struct {
	db string
}

type ProductRepositoryMongo struct {
	db string
}

// FindAll implements ProductRepository.
func (ProductRepositoryMongo) FindAll() []Product {
	panic("unimplemented")
}

// FindById implements ProductRepository.
func (ProductRepositoryMongo) FindById(id int) Product {
	panic("unimplemented")
}

// Save implements ProductRepository.
func (ProductRepositoryMongo) Save(product Product) Product {
	panic("unimplemented")
}

// FindAll implements ProductRepository.
func (ProductRepositoryPostgres) FindAll() []Product {
	panic("unimplemented")
}

// FindById implements ProductRepository.
func (ProductRepositoryPostgres) FindById(id int) Product {
	panic("unimplemented")
}

// Save implements ProductRepository.
func (ProductRepositoryPostgres) Save(product Product) Product {
	panic("unimplemented")
}

func NewProductRepository(db string) ProductRepository {
	return ProductRepositoryMongo{db: db}
}
