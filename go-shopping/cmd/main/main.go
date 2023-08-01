package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prathigajd/go-shopping/pkg/config"
	"github.com/prathigajd/go-shopping/pkg/routes"
)

func main() {
	config.Connect()
	r := mux.NewRouter()
	routes.RegisterShoppingRoutes(r)

	// Create a new HTTP server manually
	server := &http.Server{
		Addr: "127.0.0.1:8081",
		Handler: r,
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}