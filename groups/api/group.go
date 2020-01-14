package group

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MihaFriskovec/3fs-assignment/db/helpers"
	group "github.com/MihaFriskovec/3fs-assignment/groups/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var groupCollection *mongo.Collection

func init() {
	groupCollection = group.Groups()
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newGroup group.Group

	_ = json.NewDecoder(r.Body).Decode(&newGroup)

	group, err := groupCollection.InsertOne(context.TODO(), newGroup)

	if err != nil {
		log.Fatalln("Error on inserting new Group", err)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(group)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	page, err := strconv.ParseInt(r.FormValue("page"), 10, 32)
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 32)
	sort := r.FormValue("sort")
	project := r.FormValue("select")

	var groupsList []group.Group

	var query = bson.M{}

	findOptions := options.Find()

	if len(project) > 0 {
		findOptions.SetProjection(helpers.Project(project))
	}

	findOptions.SetSort(helpers.Sort(sort))

	// Pagination
	if page == 0 {
		page = 1
	}

	findOptions.SetLimit(limit)
	findOptions.SetSkip(page*limit - limit)

	list, err := groupCollection.Find(context.TODO(), query, findOptions)

	if err != nil {
		log.Fatalln("Error listing groups", err)
	}

	list.All(context.TODO(), &groupsList)

	json.NewEncoder(w).Encode(groupsList)
}

func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var group group.Group

	err := groupCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)

	if err != nil {
		log.Println("Error reading Group", err)
	}

	json.NewEncoder(w).Encode(group)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var newGroup group.Group

	_ = json.NewDecoder(r.Body).Decode(&newGroup)

	var setElements bson.D

	if len(newGroup.Name) > 0 {
		setElements = append(setElements, bson.E{"name", newGroup.Name})
	}

	updatedGroup := bson.M{"$set": setElements}

	updated, err := groupCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updatedGroup)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(updated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	deleted, err := groupCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	fmt.Println(deleted)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(deleted)
}
