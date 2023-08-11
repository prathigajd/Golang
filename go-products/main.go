package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"gothub.com/prathigajd/go-products/config"
	"gothub.com/prathigajd/go-products/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"

)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
)

func init() {
	err := cleanenv.ReadEnv(&cfg) // Pass the pointer to cfg to correctly read the environment variables.
	if err != nil {
		log.Fatal("configuration cannot be read : ", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	c, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatal("unable to connect to database : ", err)
	}

	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	h := handlers.ProductHandler{}
	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"))
	log.Printf("listening on %s:%s", cfg.Host, cfg.Port)
	log.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
