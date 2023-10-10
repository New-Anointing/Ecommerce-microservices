package product

import (
	"time"
)

type Products struct {
	Id          string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	image       []byte
	Price       float64   `json:"price"`
	Avalability bool      `json:"avalability"`
	Description string    `json:"description"`
	Unit        int32     `json:"unit"`
	createdAt   time.Time `json:"createdat"`
}
