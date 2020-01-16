package group

import (
	"log"
	"os"

	"github.com/MihaFriskovec/3fs-assignment/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Group schema
type Group struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name *string            `json:"name,omitempty" bson:"name"`
}

var collection *mongo.Collection

func init() {
	dbName := os.Getenv("DB_NAME")
	collection = db.ConnectDatabase(dbName).Collection("groups")

	if collection == nil {
		log.Println("Error connecting to users collection")
	} else {
		log.Println("Connected to users collection")
	}
}

func Groups() *mongo.Collection {
	return collection
}
