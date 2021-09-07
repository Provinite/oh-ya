package inventoryitem

import (
	"fmt"

	"github.com/provinite/oh-ya/server/serverutils"
	"gorm.io/gorm"
)

func IncrementStock(db *gorm.DB, itemId uint, quantity uint) (*InventoryItem, error) {
	inventoryItem := InventoryItem{}
	dbResult := db.First(&inventoryItem, itemId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return nil, serverutils.NotFoundError{}
	}
	inventoryItem.Stock += quantity
	dbResult = db.Model(&inventoryItem).Update("stock", inventoryItem.Stock)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &inventoryItem, nil
}
func DecrementStock(db *gorm.DB, itemId uint, quantity uint) error {
	inventoryItem := InventoryItem{}
	dbResult := db.First(&inventoryItem, itemId)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return serverutils.NotFoundError{}
	}
	if inventoryItem.Stock < quantity {
		return InsufficientStockError{
			InventoryItem: inventoryItem,
		}
	}
	inventoryItem.Stock -= quantity
	dbResult = db.Model(&inventoryItem).Update("stock", inventoryItem.Stock)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

type InsufficientStockError struct {
	InventoryItem InventoryItem
}

func (ise InsufficientStockError) Error() string {
	return fmt.Sprintf("Insufficient stock of %s", ise.InventoryItem.Name)
}
