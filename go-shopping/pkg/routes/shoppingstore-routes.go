package routes

import (
	"github.com/gorilla/mux"
	"github.com/prathigajd/go-shopping/pkg/controllers"
)

var RegisterShoppingRoutes = func(router *mux.Router) {
	router.HandleFunc("/shop/", controllers.CreateShopping).Methods("post")
	router.HandleFunc("/shop/", controllers.GetShopping).Methods("get")
  router.HandleFunc("/shop/{shoppingId}", controllers.GetShoppingByID).Methods("get")
router.HandleFunc("/shop/{shoppingId}", controllers.UpdateShopping).Methods("put")
router.HandleFunc("/shop/{shoppingID}", controllers.DeleteShopping).Methods("delete")
}