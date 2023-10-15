package services

import (
	"sesi-6/models"

	"gorm.io/gorm"
)

func CreateVariant(db *gorm.DB, variant *models.Variant) error {
	return db.Create(&variant).Error
}

func UpdateVariantById(db *gorm.DB, id uint, variant *models.Variant) error {
	return db.Model(&models.Variant{}).Where("id = ?", id).Updates(variant).Error
}

func DeleteVariantById(db *gorm.DB, id uint) error {
	return db.Delete(&models.Variant{}, id).Error
}
