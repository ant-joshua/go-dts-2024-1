package repository

import "sesi_11/entity"

type ProductRepository interface {
	FindById(id int) *entity.Product
}
