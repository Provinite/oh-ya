package recipe

import (
	"github.com/provinite/oh-ya/inventoryitem"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model      `json:"-"`
	ID              uint                        `json:"id"`
	InventoryItemID uint                        `json:"output_item"`
	InventoryItem   inventoryitem.InventoryItem `json:"-"`
	Quantity        uint                        `json:"quantity"`
	RecipeInputs    []RecipeInput               `json:"inputs"`
}
