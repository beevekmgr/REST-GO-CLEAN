package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ProductUID string `json:"user_uid"`
	Name       string `json:"name"`
	Price      string `json:"price"`
}

// Migrate product
func (c Product) Migrate(db *gorm.DB) error {
	if m := db.AutoMigrate(c); m.Error != nil {
		return m.Error
	}
	return nil
}

// ToMap converts member to map
func (c Product) ToMap() gin.H {
	return gin.H{
		"id":    c.ID,
		"name":  c.Name,
		"price": c.Price,
	}
}
