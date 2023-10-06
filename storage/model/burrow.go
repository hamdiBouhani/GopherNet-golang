package model

import (
	"time"

	"github.com/google/uuid"
)

type Burrow struct {
	UUID uuid.UUID `json:"uuid" gorm:"type:uuid;"`
	ID   int64     `json:"id" gorm:"primary_key;autoincrement"`

	Name     string  `json:"name" gorm:"column:name"`
	Depth    float64 `json:"depth" gorm:"column:depth"`
	Wide     float64 `json:"wide" gorm:"column:wide"`
	Occupied bool    `json:"occupied" gorm:"column:occupied"`
	Age      int     `json:"age" gorm:"column:age"`

	CreatedAt time.Time  `json:"created_date" gorm:"column:created_date"`
	UpdatedAt time.Time  `json:"updated_date" gorm:"column:changed_date"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_date"`
}

func (Burrow) TableName() string {
	return "burrow"
}
