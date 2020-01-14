package group

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MihaFriskovec/3fs-assignment/db/helpers"
	"github.com/MihaFriskovec/3fs-assignment/errors"
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
	var newGroup group.Group

	json.NewDecoder(r.Body).Decode(&newGroup)

	group, err := groupCollection.InsertOne(context.TODO(), newGroup)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error creating Group", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error creating Group", 500))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func List(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Fatalln("Error listing Groups", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error listing Groups", 500))
		return
	}

	list.All(context.TODO(), &groupsList)
	json.NewEncoder(w).Encode(groupsList)
}

func Read(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var group group.Group

	err := groupCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error reading Group", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error reading Group", 500))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var newGroup group.Group

	json.NewDecoder(r.Body).Decode(&newGroup)

	var setElements bson.D

	if len(newGroup.Name) > 0 {
		setElements = append(setElements, bson.E{"name", newGroup.Name})
	}

	updatedGroup := bson.M{"$set": setElements}

	updated, err := groupCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updatedGroup)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error updating Group", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error updating Group", 500))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	deleted, err := groupCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error deleting Group", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error deleting Group", 500))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleted)
}
