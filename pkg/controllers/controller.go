package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"catalogService/pkg/models"
	"catalogService/pkg/utils"
	"github.com/gorilla/mux"
)

var Product models.Catalog

func CreateProduct(w http.ResponseWriter, r *http.Request)  {
	 CreateProduct := &models.Catalog{}
	 utils.ParseBody(r, CreateProduct)
	 product_obj := CreateProduct.CreateProduct()
	 res,_ := json.Marshal(product_obj)
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)

}

func GetProduct(w http.ResponseWriter, r *http.Request)  {
	products := models.GetAllProduct()
	res,_ :=json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	product_id := vars["pid"]
	ID,err := strconv.ParseInt(product_id,0,0)
	if err != nil {
		fmt.Println("Error while parsing!!!!")
	}
	productDetails,_ := models.GetProductById(ID)
	res,_ := json.Marshal(productDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PutProduct(w http.ResponseWriter, r *http.Request)  {
	var updateProduct = &models.Catalog{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	productid := vars["pid"]
	ID, err := strconv.ParseInt(productid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := models.GetProductById(ID)
	if productDetails.Product_name != "" {
		productDetails.Product_name = updateProduct.Product_name
	}

	db.Save(&productDetails)
	res, _ := json.Marshal(productDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	product_id := vars["pid"]
	ID, err := strconv.ParseInt(product_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}