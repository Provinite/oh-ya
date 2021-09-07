package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/provinite/oh-ya/batch"
	"github.com/provinite/oh-ya/db"
	"github.com/provinite/oh-ya/inventoryitem"
	"github.com/provinite/oh-ya/purchase"
	"github.com/provinite/oh-ya/recipe"
	"github.com/provinite/oh-ya/sale"
)

func StartServer() error {
	if db.DBError != nil {
		log.Fatal(db.DBError.Error())
	}
	db.DB.AutoMigrate(&inventoryitem.InventoryItem{})
	db.DB.AutoMigrate(&purchase.Purchase{})
	db.DB.AutoMigrate(&purchase.PurchaseLine{})
	db.DB.AutoMigrate(&sale.Sale{})
	db.DB.AutoMigrate(&sale.SaleLine{})
	db.DB.AutoMigrate(&recipe.Recipe{})
	db.DB.AutoMigrate(&recipe.RecipeInput{})
	db.DB.AutoMigrate(&batch.Batch{})

	router := httprouter.New()

	router.POST("/inventory-items", inventoryitem.ListPost)
	router.GET("/inventory-items", inventoryitem.ListGet)
	router.GET("/inventory-items/:inventoryitem", inventoryitem.DetailGet)

	router.POST("/purchases", purchase.ListPost)
	router.GET("/purchases", purchase.ListGet)

	router.POST("/sales", sale.ListPost)
	router.GET("/sales", sale.ListGet)

	router.POST("/recipes", recipe.ListPost)
	router.GET("/recipes", recipe.ListGet)

	router.POST("/batches", batch.ListPost)
	router.GET("/batches", batch.ListGet)
	return http.ListenAndServe(":8080", Cors(ContentType(router)))
}

func ContentType(handler http.Handler) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(ourFunc)
}

func Cors(handler http.Handler) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("access-control-allow-origin", r.Header.Get("Origin"))
		w.Header().Add("access-control-allow-headers", "*")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(ourFunc)
}
