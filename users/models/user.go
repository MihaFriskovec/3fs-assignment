package user

import (
	"log"

	"github.com/MihaFriskovec/3fs-assignment/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User schema
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Email    *string            `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Group    primitive.ObjectID `json:"group,omitempty" bson:"group,omitempty"`
}

var collection *mongo.Collection

func init() {
	collection = db.ConnectDatabase("3fs").Collection("users")

	if collection == nil {
		log.Println("DB ERROR: Error connecting to users collection")
	} else {
		log.Println("Connected to users collection")
	}
}

func Users() *mongo.Collection {
	return collection
}

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Println("Error Hasisg password", err)
	}

	return string(hash)
}
