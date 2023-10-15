package services

import (
	"sesi-6/models"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product *models.Product) error {
	return db.Create(&product).Error
}

func UpdateProduct(db *gorm.DB, id uint, product *models.Product) error {
	return db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func GetProductById(db *gorm.DB, id uint) (*models.Product, error) {
	var product models.Product
	err := db.First(&product, id).Error
	return &product, err
}

func GetProductWithVariant(db *gorm.DB, id uint) (*models.Product, error) {
	var product models.Product
	err := db.Preload("Variants").First(&product, id).Error
	return &product, err
}
