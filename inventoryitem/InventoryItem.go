package inventoryitem

import (
	"time"

	"gorm.io/gorm"
)

type InventoryItem struct {
	gorm.Model `json:"-"`
	ID         uint `json:"id"`
	// Name of the inventory item
	Name string `json:"name"`
	// Cost in cents to aquire, if this inventory item can
	// be purchased directly
	Cost uint `json:"cost_cents"`
	// Price in cents to sell, if this inventory item can
	// be sold
	Price uint `json:"price_cents"`
	// ForSale flag true if inventory item is for sale
	ForSale bool `json:"for_sale"`
	// CanPurchase flag true if inventory item can be purchased
	CanPurchase bool           `json:"can_purchase"`
	Stock       uint           `json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
