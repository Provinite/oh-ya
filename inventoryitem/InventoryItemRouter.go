package inventoryitem

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/server/serverutils"
	"gorm.io/gorm"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serverutils.ListGet(w, req, params, &[]InventoryItem{})
}

func DetailGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	idStr := params.ByName("inventoryitem")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		inventoryItem := InventoryItem{}
		rowsAffected := db.DB.Find(&inventoryItem, id).RowsAffected
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(inventoryItem)
		}
	}
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	body := &struct {
		Name  string `json:"name"`
		Cost  *uint  `json:"cost_cents"`
		Price *uint  `json:"price_cents"`
		Stock *uint  `json:"stock"`
	}{}
	json.NewDecoder(req.Body).Decode(body)
	cost := uint(0)
	price := uint(0)
	stock := uint(0)
	canBuy := false
	canSell := false
	if body.Cost != nil {
		cost = *body.Cost
		canBuy = true
	}
	if body.Price != nil {
		price = *body.Price
		canSell = true
	}
	if body.Stock != nil {
		stock = *body.Stock
	}
	inventoryItem := &InventoryItem{
		Name:        body.Name,
		Cost:        cost,
		Price:       price,
		ForSale:     canSell,
		CanPurchase: canBuy,
		DeletedAt:   gorm.DeletedAt{},
		Stock:       stock,
	}
	db.DB.Create(inventoryItem)
	json.NewEncoder(w).Encode(inventoryItem)
}
