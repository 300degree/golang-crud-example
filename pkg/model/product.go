package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        uuid.UUID `gorm:"type:uuid; primary_key; column:id; default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"type:varchar; column:name"`
	Price     int64     `json:"price" gorm:"column:price"`
	CreateAt  time.Time `gorm:"autoCreateTime; column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime; column:updated_at"`
}
