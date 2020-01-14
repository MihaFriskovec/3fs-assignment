package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Group schema
type Group struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name"`
}

// User schema
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Group    primitive.ObjectID `json:"group" bson:"groups"`
}

var client *mongo.Client

func initDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://appUser:byhvo2-rigDiq-jincam@3fs-bhpqi.mongodb.net/3fs?retryWrites=true&w=majority")

	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	err1 := client.Ping(context.Background(), readpref.Primary())

	if err1 != nil {
		log.Fatal("Couldn't ping to the database", err1)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
}

func listGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	page, err := strconv.ParseInt(r.FormValue("page"), 10, 32)
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 32)
	sort := r.FormValue("sort")
	project := r.FormValue("select")

	// Projection
	projectFields := strings.Split(project, ",")

	projectionQuery := bson.M{}

	for _, f := range projectFields {
		projectionQuery[f] = 1
	}

	fmt.Println(projectionQuery)
	// END Projection

	// SORT
	var sortFlied string

	sortOrder := sort[0:1]
	var sortOrderQuery int

	if sortOrder == "-" || sortOrder == "+" {
		sortFlied = sort[1:]
		if sortOrder == "-" {
			sortOrderQuery = -1
		}
	} else {
		sortOrderQuery = 1
		sortFlied = sort
	}
	// END SORT

	var groups []Group
	groupsCollection := client.Database("3fs").Collection("groups")

	var query = bson.M{}

	fmt.Println(projectionQuery)

	findOptions := options.Find()
	// Pagination
	findOptions.SetLimit(limit)
	findOptions.SetSkip(page*limit - limit)
	// Sort
	findOptions.SetSort(bson.M{sortFlied: sortOrderQuery})
	// Project/Select
	if len(projectFields) > 0 {
		findOptions.SetProjection(projectionQuery)
	}

	list, err := groupsCollection.Find(context.TODO(), query, findOptions)

	if err != nil {
		log.Fatalln("Error", err)
	}

	list.All(context.TODO(), &groups)

	json.NewEncoder(w).Encode(groups)
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	groupsCollection := client.Database("3fs").Collection("groups")

	var newGroup Group

	_ = json.NewDecoder(r.Body).Decode(&newGroup)

	group, err := groupsCollection.InsertOne(context.TODO(), newGroup)

	if err != nil {
		log.Fatalln("Error on inserting new Group", err)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(group)
}

func readGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	groupsCollection := client.Database("3fs").Collection("groups")

	var group Group

	err := groupsCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)

	if err != nil {
		log.Fatalln("Error reading Group", err)
	}

	json.NewEncoder(w).Encode(group)
}

func createGroups(w http.ResponseWriter, r *http.Request) {

}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []User
	usersCollection := client.Database("3fs").Collection("users")

	list, err := usersCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatalln("Error listing User", err)
	}

	list.All(context.TODO(), &users)

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	usersCollection := client.Database("3fs").Collection("users")

	var newUser User

	_ = json.NewDecoder(r.Body).Decode(&newUser)

	group, err := usersCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Fatalln("Error on inserting new User", err)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(group)
}

func readUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	usersCollection := client.Database("3fs").Collection("users")

	var user User

	err := usersCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		log.Fatalln("Error reading User", err)
	}

	json.NewEncoder(w).Encode(user)
}

func server() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/groups", listGroups).Methods("GET")
	router.HandleFunc("/api/groups", createGroup).Methods("POST")
	router.HandleFunc("/api/groups/bulk", createGroups).Methods("POST")
	router.HandleFunc("/api/groups/{id}", readGroup).Methods("GET")

	router.HandleFunc("/api/users", listUsers).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", readUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
