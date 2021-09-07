package batch

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/inventoryitem"
	"github.com/provinite/oh-ya/recipe"
	"github.com/provinite/oh-ya/server/serverutils"
	"gorm.io/gorm"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	batches := []Batch{}
	serverutils.ListGet(w, req, params, &batches)
}

func DetailGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	body := struct {
		RecipeID *uint `json:"recipe"`
		Quantity *uint `json:"quantity"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.RecipeID == nil {
		http.Error(w, "Bad Request", 400)
	}

	err = db.DB.Transaction(func(db *gorm.DB) error {
		recipe := recipe.Recipe{}
		result := db.Preload("RecipeInputs").First(&recipe, *body.RecipeID)
		if result.RowsAffected == 0 {
			return serverutils.BadRequestError{}
		}
		if result.Error != nil {
			return result.Error
		}

		// create a batch
		batch := Batch{
			RecipeID: recipe.ID,
		}

		if body.Quantity != nil {
			batch.Quantity = *body.Quantity
		} else {
			batch.Quantity = 1
		}
		result = db.Create(&batch)
		if result.Error != nil {
			return result.Error
		}

		// decrement inputs
		var errors []error
		for _, input := range recipe.RecipeInputs {
			err := inventoryitem.DecrementStock(db, input.InventoryItemID, input.Quantity*batch.Quantity)
			if err != nil {
				errors = append(errors, err)
			}
		}

		if len(errors) > 0 {
			return serverutils.CompoundError{
				Errors: errors,
			}
		}

		// increment output
		_, err := inventoryitem.IncrementStock(db, recipe.InventoryItemID, recipe.Quantity*batch.Quantity)
		if err != nil {
			return err
		}

		json.NewEncoder(w).Encode(&batch)
		return nil
	})
	compoundError, ok := err.(serverutils.CompoundError)
	errors := []error{}
	if ok {
		errors = append(errors, compoundError.Errors...)
	} else if err != nil {
		errors = append(errors, err)
	}
	code := 200
	var message string
	if err != nil {
		message = err.Error()
	}
	for _, err := range errors {
		if err != nil {
			switch err.(type) {
			case inventoryitem.InsufficientStockError:
				{
					if code < 400 {
						code = 400
					}
				}
			case serverutils.BadRequestError:
				{
					if code < 400 {
						code = 400
					}
				}
			default:
				{
					code = 500
				}
			}
		}
	}
	if code != 200 {
		http.Error(w, message, code)
	}
}
