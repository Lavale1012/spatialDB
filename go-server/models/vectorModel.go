package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type VectorModel struct {
	gorm.Model
	ID     string          `gorm:"primaryKey;uniqueIndex;not null" json:"id"`
	Vector pq.Float64Array `gorm:"type:float8[]"`
	Text   string          `gorm:"not null" json:"text"` // Added Text field to store the original text
}
