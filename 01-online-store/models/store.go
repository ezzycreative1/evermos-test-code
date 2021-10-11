package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int64          `gorm:"primaryKey";"AUTO_INCREMENT" json:"id"`
	Name        string         `gorm:"null" json:"name"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	Kota        string         `json:"kota"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" json:"deleted_at"`
}

type Category struct {
	ID        int64          `gorm:"primaryKey";"AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"null" json:"name"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" json:"deleted_at"`
}

type Product struct {
	ID            int64          `gorm:"primaryKey";"AUTO_INCREMENT" json:"id"`
	CategoryID    int64          `gorm:"not null;index" json:"category_id"`
	SupplierID    int64          `gorm:"not null;index" json:"supplier_id"`
	Name          string         `gorm:"null" json:"name"`
	Description   string         `json:"description"`
	Unit          string         `json:"unit"`
	Size          string         `json:"size"`
	Content       string         `json:"content"`
	PurchasePrice int64          `json:"purchase_price"`
	SellingPrice  int64          `json:"selling_price"`
	Stock         int64          `json:"stock"`
	Status        string         `json:"status"`
	EntryDate     time.Time      `gorm:"not null" json:"entry_date"`
	CreatedAt     time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" json:"deleted_at"`
}

type Supplier struct {
	ID          int64          `gorm:"primaryKey";"AUTO_INCREMENT" json:"id"`
	Name        string         `gorm:"null" json:"name"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	Kota        string         `json:"kota"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" json:"deleted_at"`
}
type Order struct {
	ID          int64          `gorm:"primaryKey";"AUTO_INCREMENT" json:"id"`
	CustomerID  int64          `gorm:"not null;index" json:"category_id"`
	Email       string         `gorm:"not null,unique" json:"email,omitempty"`
	Address     string         `json:"address"`
	PhoneNumber string         `json:"phone_number"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" json:"deleted_at"`
}

type OrderDetail struct {
	OrderID   int64 `gorm:"not null;index" json:"order_id"`
	ProductID int64 `gorm:"not null;index" json:"product_id"`
	Amount    int64 `json:"amount"`
}
