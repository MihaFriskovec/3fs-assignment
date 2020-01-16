package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbType := os.Getenv("DB_TYPE")
	dbPrefix := os.Getenv("DB_PREFIX")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	var connectionURL string

	if dbType == "local" {
		connectionURL = fmt.Sprintf("%s://%s/%s", dbPrefix, dbURL, dbName)
	} else {
		connectionURL = fmt.Sprintf("%s://%s:%s@%s/%s?retryWrites=true&w=majority", dbPrefix, dbUser, dbPassword, dbURL, dbName)
	}

	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	c = client
	log.Println("Connected to MongoDB")

	// err1 := client.Ping(context.Background(), readpref.Primary())

	// if err1 != nil {
	// 	log.Fatal("Couldn't ping to the database", err1)
	// } else {
	// 	log.Println("Connected to MongoDB")

	// }
}

func ConnectDatabase(databaseName string) *mongo.Database {
	database := c.Database(databaseName)

	log.Println("Connected to Database: ", databaseName)

	return database
}
