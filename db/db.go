package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var c *mongo.Client

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectionURL := fmt.Sprintf("mongodb+srv://%s:%s@3fs-bhpqi.mongodb.net/%s?retryWrites=true&w=majority", dbUser, dbPassword, dbName)

	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	err1 := client.Ping(context.Background(), readpref.Primary())

	if err1 != nil {
		log.Fatal("Couldn't ping to the database", err1)
	} else {
		log.Println("Connected to MongoDB")

		c = client
	}
}

func ConnectDatabase(databaseName string) *mongo.Database {
	database := c.Database(databaseName)

	log.Println("Connected to Database: ", databaseName)

	return database
}
