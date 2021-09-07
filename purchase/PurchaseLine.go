package purchase

import (
	"github.com/provinite/oh-ya/inventoryitem"
	"gorm.io/gorm"
)

type PurchaseLine struct {
	gorm.Model      `json:"-"`
	ID              uint                        `json:"id" gorm:"primaryKey"`
	InventoryItemID uint                        `json:"inventory_item"`
	InventoryItem   inventoryitem.InventoryItem `json:"-"`
	PurchaseID      uint                        `json:"purchase"`
	Cost            uint                        `json:"cost_cents"`
	Quantity        uint                        `json:"quantity"`
}
