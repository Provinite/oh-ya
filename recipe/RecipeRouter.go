package recipe

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/server/serverutils"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	dest := []Recipe{}
	skip, limit, _ := serverutils.GetPagination(req)
	result := db.DB.Preload("RecipeInputs").Limit(limit).Offset(skip).Find(&dest)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
	json.NewEncoder(w).Encode(dest)
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	body := struct {
		InventoryItemID uint `json:"output_item"`
		Quantity        uint `json:"quantity"`
		RecipeInputs    []struct {
			InventoryItemID uint `json:"inventory_item"`
			Quantity        uint `json:"quantity"`
		} `json:"inputs"`
	}{}
	json.NewDecoder(req.Body).Decode(&body)
	recipe := Recipe{
		InventoryItemID: body.InventoryItemID,
		Quantity:        body.Quantity,
	}
	for _, pl := range body.RecipeInputs {
		appendResult := append(recipe.RecipeInputs, RecipeInput{
			InventoryItemID: pl.InventoryItemID,
			Quantity:        pl.Quantity,
		})
		recipe.RecipeInputs = appendResult
	}
	db.DB.Create(&recipe)
	json.NewEncoder(w).Encode(&recipe)
}
