package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int64          `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"null" json:"name"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	Password    string         `gorm:"not null" json:"password,omitempty"`
	PhoneNumber string         `json:"phone_number"`
	Photo       string         `json:"photo"`
	Address     string         `json:"address"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:",omitempty"`
}
