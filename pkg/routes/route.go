package routes

import (
	"catalogService/pkg/controllers"
	"github.com/gorilla/mux"
)

var CatalogService =  func(router *mux.Router)  {
	router.HandleFunc("/catalog/product",controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/catalog/product",controllers.GetProduct).Methods("GET")
	router.HandleFunc("/catalog/product/{pid}",controllers.GetProductById).Methods("GET")
	router.HandleFunc("/catalog/product/{pid}",controllers.PutProduct).Methods("PUT")
	router.HandleFunc("/catalog/product/{pid}",controllers.DeleteProduct).Methods("DELETE")

}