package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int       `json:"id" gorm:"column:product_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:product_name"`
	Price       int       `json:"price" gorm:"column:product_price"`
	RetailPrice int       `json:"retail_price" gorm:"column:retail_price"`
	CategoryID  int       `json:"category_id" gorm:"column:category_id;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	Category    *Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	// margin 10%
	p.RetailPrice = p.Price + (p.Price * 10 / 100)
	return
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	// margin 10%
	p.RetailPrice = p.Price + (p.Price * 10 / 100)
	return
}

type GetListProductRequest struct {
	Page       int     `json:"page" form:"page"`
	Limit      int     `json:"limit" form:"limit"`
	Name       *string `json:"name" form:"name"`
	StartPrice *int    `json:"start_price" form:"start_price"`
	EndPrice   *int    `json:"end_price" form:"end_price"`
}

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type UpdateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
