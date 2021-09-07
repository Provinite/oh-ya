package serverutils

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/db"
)

func ListGet(w http.ResponseWriter, req *http.Request, params httprouter.Params, dest interface{}) {
	skip, limit, _ := GetPagination(req)
	db.DB.Limit(limit).Offset(skip).Find(dest)
	json.NewEncoder(w).Encode(dest)
}
