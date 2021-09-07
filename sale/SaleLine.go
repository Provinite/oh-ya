package sale

import (
	"time"

	"github.com/provinite/oh-ya/inventoryitem"
	"gorm.io/gorm"
)

type SaleLine struct {
	gorm.Model      `json:"-"`
	ID              uint                        `json:"id"`
	InventoryItemID uint                        `json:"inventory_item"`
	InventoryItem   inventoryitem.InventoryItem `json:"-"`
	SaleID          uint                        `json:"-"`
	Price           uint                        `json:"price_cents"`
	Cost            uint                        `json:"cost_cents"`
	Quantity        uint                        `json:"quantity"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
	DeletedAt       gorm.DeletedAt              `gorm:"index" json:"deleted_at"`
}
