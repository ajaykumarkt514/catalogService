package main

import (
	"github.com/gorilla/mux"
	 "catalogService/pkg/routes"
	"net/http"
)

func main()  {
	router := mux.NewRouter()
	routes.CatalogService(router)
	http.ListenAndServe("localhost:8024",router)
}
