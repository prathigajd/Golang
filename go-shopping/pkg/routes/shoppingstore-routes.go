package routes

import (
	"github.com/gorilla/mux"
	"github.com/prathigajd/go-shopping/pkg/controllers"
)

func RegisterShoppingRoutes(router *mux.Router) {
	// Define routes for creating, getting, updating, and deleting shopping items
	router.HandleFunc("/shop", controllers.CreateShopping).Methods("POST")
	router.HandleFunc("/shop", controllers.GetShopping).Methods("GET")
	router.HandleFunc("/shop/{shoppingId}", controllers.GetShoppingByID).Methods("GET")
	router.HandleFunc("/shop/{shoppingId}", controllers.UpdateShopping).Methods("PUT")
	router.HandleFunc("/shop/{shoppingId}", controllers.DeleteShopping).Methods("DELETE")
}
