package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"gothub.com/prathigajd/go-products/dbiface"
)

var (
	v = validator.New()
)

type ProductHandler struct {
	col dbiface.CollectionAPI // Replace with your actual collection interface
}

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

type product struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"product_name" bson:"product_name" validate:"required,max=10"`
	Price    int                `json:"price" bson:"price" validate:"required,max=2000"`
	Currency string             `json:"currency" bson:"currency" validate:"required,len=3"`
	Discount int                `json:"discount" bson:"discount"`
}

func createProducts(ctx context.Context, col dbiface.CollectionAPI, products []product) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, product := range products {
		product.ID = primitive.NewObjectID()
		insertID, err := col.InsertOne(ctx, product) // Use col.InsertOne instead of collection.InsertOne
		if err != nil {
			log.Printf("unable to insert: %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func (h *ProductHandler) CreateProducts(c echo.Context) error {
	var products []product
	c.Echo().Validator = &ProductValidator{validator: v}
	if err := c.Bind(&products); err != nil {
		log.Printf("unable to bind: %v", err)
		return err
	}
	for _, product := range products {
		if err := c.Validate(product); err != nil {
			log.Printf("unable to validate the product %+v %v ", product, err)
			return err
		}
	}
	IDs, err := createProducts(context.Background(), h.col, products)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs) // Return the inserted IDs as JSON response
}

/*func insertProducts(ctx context.Context, products []product, invalid type) {
	err := insertProducts(ctx, products)
	if err != nil {
		panic(err)
}
}*/
