package sale

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/inventoryitem"
	"github.com/provinite/oh-ya/recipe"
	"github.com/provinite/oh-ya/server/serverutils"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	dest := []Sale{}
	skip, limit, _ := serverutils.GetPagination(req)
	result := db.DB.Preload("SaleLines").Limit(limit).Offset(skip).Find(&dest)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
	json.NewEncoder(w).Encode(dest)
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	body := struct {
		SaleLines []struct {
			InventoryItemID uint `json:"inventory_item"`
			Quantity        uint `json:"quantity"`
		} `json:"sale_lines"`
	}{}
	json.NewDecoder(req.Body).Decode(&body)
	sale := Sale{}

	for _, pl := range body.SaleLines {
		inventoryItem := inventoryitem.InventoryItem{}
		db.DB.First(&inventoryItem, pl.InventoryItemID)
		cost := uint(0)
		if !inventoryItem.CanPurchase {
			recipe := recipe.Recipe{}
			result := db.DB.Preload("RecipeInputs").Preload("RecipeInputs.InventoryItem").First(&recipe, "inventory_item_id = ?", inventoryItem.ID)
			if result.Error != nil {
				http.Error(w, result.Error.Error(), 500)
				return
			}
			if result.RowsAffected != 0 {
				for _, recipeInput := range recipe.RecipeInputs {
					cost += recipeInput.InventoryItem.Cost * recipeInput.Quantity
				}
				cost = cost / recipe.Quantity
			}
		} else {
			cost = inventoryItem.Cost
		}
		saleLine := SaleLine{
			InventoryItemID: pl.InventoryItemID,
			Quantity:        pl.Quantity,
			Price:           inventoryItem.Price,
			Cost:            cost,
		}
		sale.SaleLines = append(sale.SaleLines, saleLine)
	}
	db.DB.Create(&sale)
	json.NewEncoder(w).Encode(&sale)
}
