package models

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Variants  []Variant `gorm:"foreignKey:ProductID"`
}
