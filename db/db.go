package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var c *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://appUser:byhvo2-rigDiq-jincam@3fs-bhpqi.mongodb.net/3fs?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	err1 := client.Ping(context.Background(), readpref.Primary())

	if err1 != nil {
		log.Fatal("Couldn't ping to the database", err1)
	} else {
		fmt.Println("Connected to MongoDB!")

		c = client
	}
}

func ConnectDatabase(databaseName string) *mongo.Database {
	database := c.Database(databaseName)

	fmt.Println("Connected to Database!", databaseName)

	return database
}
