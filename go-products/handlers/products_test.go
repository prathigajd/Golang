package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stretchr/testify/assert"
	"gothub.com/prathigajd/go-products/config"

	//"gothub.com/prathigajd/go-products/handlers"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
	h *ProductHandler
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

func TestProduct(t *testing.T) {
	t.Run("test create priduct", func(t *testing.T) {
		body := `
	   [{
			"product_name": "googletalk",
			"price": 250,
			"currency":"USD",
			"vendor":"google",
			"accessories":["charger", "subscription"]
		}] `
		
		req := httptest.NewRequest("POST", "/products", strings.NewReader(body))
	    res := httptest.NewRecorder()
		//req.Header
		e := echo.New()
		c:= e.NewContext(req, res)
		h.col = col
		err := h.CreateProducts(c)
		assert.Nil(t, err)
		

	})
}