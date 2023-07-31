package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
    "github.com/prathigajd/go-shopping/pkg/utils"
	"github.com/prathigajd/go-shopping/pkg/models"
	
)

var NewShopping models.Shopping

func GetShopping(w http.ResponseWriter, r *http.Request){
	newShoppings := models.GetAllShoppings()
	res, _ := json.Marshal(newShoppings)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetShoppingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ShoppingID := vars["shoppingID"]
	ID, err := strconv.ParseInt(ShoppingID, 0, 0)

	if err != nil {
		fmt.Println("error whie parsing")
	}
	ShoppingDetails, _ := models.GetShoppingByID(ID)
	res, _ := json.Marshal(ShoppingDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
} 

func CreateShopping(w http.ResponseWriter, r *http.Request) {
CreateShopping := &models.Shopping{}
utils.ParseBody(r, CreateShopping)
b := CreateShopping.Createshopping()
res, _ := json.Marshal(b)
w.WriteHeader(http.StatusOK)
w.Write(res)
}

func DeleteShopping(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	shoppingID := vars["shoppingID"]
	ID, err := strconv.ParseInt(shoppingID, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	shopping := models.DeleteShopping(ID)
	res, _ := json.Marshal(shopping)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateShopping(w http.ResponseWriter, r *http.Request) {
	var updateShopping = &models.Shopping{}
	utils.ParseBody(r, updateShopping)
	vars := mux.Vars(r)
	ShoppingId := vars["shoppingId"]
    ID, err := strconv.ParseInt(ShoppingId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
ShoppingDetails, db := models.GetShoppingByID(ID)
if updateShopping.Snumber != "" {
	ShoppingDetails.Snumber = updateShopping.Snumber
}
if updateShopping.Items != "" {
	ShoppingDetails.Items = updateShopping.Items
}
if updateShopping.Price != "" {
	ShoppingDetails.Price = updateShopping.Price
}
db.Save(&ShoppingDetails)
res, _ := json.Marshal(ShoppingDetails)
w.Header().Set("content-Type", "pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}