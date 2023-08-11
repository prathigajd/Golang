package dbiface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionAPI interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

// package dbiface

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type CollectionAPI interface {
// 	InsertOne(ctx context.Context, document interface{}) (*primitive.ObjectID, error)
// 	// Define other methods for collection interaction (e.g., Find, Update, Delete, etc.)
// }
