package ingredient

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Header().Add("Content-type", "application/json")
	slice := make([]Ingredient, 3)
	slice[0] = Ingredient{Name: "Beebo", ID: 100}
	slice[1] = Ingredient{Name: "Baggins", ID: 101}
	slice[2] = Ingredient{Name: "AndSam", ID: 102}
	json.NewEncoder(w).Encode(slice)
}

func DetailGet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Header().Add("Content-type", "application/json")
	id, err := strconv.ParseUint(params.ByName("ingredient"), 10, 64)
	if err != nil {
		// TODO
		return
	}
	json.NewEncoder(w).Encode(Ingredient{
		Name: "Ingredient Name",
		ID:   id,
	})
}

func ListPost(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Header().Add("Content-type", "application/json")
	body := &struct {
		Name string `json:"name"`
	}{}
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		// TODO
		return
	}
	ingredient := &Ingredient{
		Name: body.Name,
	}
	result := db.DB.Create(ingredient)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
	json.NewEncoder(w).Encode(ingredient)
}
