package models

// ProductRecord struct models, collection of record product management
type Product struct {
	Model
	ProductName        string `gorm:"column:product_name"`
	ProductDescription string `gorm:"column:product_description"`
	Quantity           int    `gorm:"column:quantity"`
}
