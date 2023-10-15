package models

import (
	"time"
)

type Variant struct {
	ID          uint      `gorm:"primaryKey"`
	VariantName string    `json:"variant_name" gorm:"type:varchar(255);not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	ProductID   uint      `json:"product_id" gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
