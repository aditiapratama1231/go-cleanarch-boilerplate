package entities

import (
	base "product-microservice/domain/infrastructure"
)

// ProductRecord struct models, collection of record product management
type Product struct {
	base.Model
	ProductName        string `gorm:"column:product_name" json:"product_name"`
	ProductDescription string `gorm:"column:product_description" json:"product_description"`
	Quantity           int    `gorm:"column:quantity" json:"quantity"`
}
