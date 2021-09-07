package purchase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/inventoryitem"
	"github.com/provinite/oh-ya/server/serverutils"
	"gorm.io/gorm"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	dest := []Purchase{}
	skip, limit, _ := serverutils.GetPagination(req)
	result := db.DB.Preload("PurchaseLines").Limit(limit).Offset(skip).Find(&dest)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
	json.NewEncoder(w).Encode(dest)
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := db.DB.Transaction(func(db *gorm.DB) error {
		body := struct {
			PurchaseLines []struct {
				InventoryItemID uint `json:"inventory_item"`
				Quantity        uint `json:"quantity"`
				Cost            uint `json:"cost_cents"`
			} `json:"purchase_lines"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			return err
		}

		purchase := Purchase{
			DeletedAt: gorm.DeletedAt{},
		}

		for _, pl := range body.PurchaseLines {
			newPurchaseLine := PurchaseLine{
				InventoryItemID: pl.InventoryItemID,
				Quantity:        pl.Quantity,
			}
			purchase.PurchaseLines = append(purchase.PurchaseLines, newPurchaseLine)
			inventoryItem, err := inventoryitem.IncrementStock(db, pl.InventoryItemID, pl.Quantity)
			newPurchaseLine.Cost = inventoryItem.Cost
			if err != nil {
				return err
			}
		}

		err = db.Create(&purchase).Error
		if err != nil {
			return err
		}

		err = json.NewEncoder(w).Encode(&purchase)
		return err
	})
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}
