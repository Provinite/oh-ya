package recipe

import (
	"github.com/provinite/oh-ya/inventoryitem"
	"gorm.io/gorm"
)

type RecipeInput struct {
	gorm.Model      `json:"-"`
	ID              uint `json:"id"`
	InventoryItemID uint `json:"inventory_item"`
	InventoryItem   inventoryitem.InventoryItem
	RecipeID        uint `json:"-"`
	Quantity        uint `json:"quantity"`
}
