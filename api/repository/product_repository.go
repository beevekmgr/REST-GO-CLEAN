package repository

import (
	"REST-GO-CLEAN/lib"
)

type ProductRepository struct {
	db lib.Database
}

// NewProductRepository creates product repository
func NewProductRepository(db lib.Database) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

//  GetProducts present in database
// func (c ProductRepository) GetProducts() (members []models.Product, err error) {
// 	return members, c.db.DB.Where("role LIKE ?", "%admin%").Find(&members).Error
// }
