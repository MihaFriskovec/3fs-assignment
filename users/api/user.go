package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MihaFriskovec/3fs-assignment/db/helpers"
	"github.com/MihaFriskovec/3fs-assignment/errors"
	user "github.com/MihaFriskovec/3fs-assignment/users/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func init() {
	userCollection = user.Users()
}

func Create(w http.ResponseWriter, r *http.Request) {
	var newUser user.User

	json.NewDecoder(r.Body).Decode(&newUser)

	passwordHash := user.HashPassword([]byte(newUser.Password))

	newUser.Password = passwordHash

	user, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Println("Error creating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error creating new user", 500))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func List(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseInt(r.FormValue("page"), 10, 32)
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 32)
	sort := r.FormValue("sort")
	project := r.FormValue("select")

	var usersList []user.User

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

	list, err := userCollection.Find(context.TODO(), query, findOptions)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		log.Println("Error listing users", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error reading users", 500))
		return
	}

	list.All(context.TODO(), &usersList)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersList)
}

func Read(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var user user.User

	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error reading user", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "User with given id not found", 404))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var newUser user.User

	json.NewDecoder(r.Body).Decode(&newUser)

	var setElements bson.D

	if len(newUser.Name) > 0 {
		setElements = append(setElements, bson.E{"name", newUser.Name})
	}

	if len(newUser.Email) > 0 {
		setElements = append(setElements, bson.E{"email", newUser.Email})
	}

	if len(newUser.Group) > 0 {
		setElements = append(setElements, bson.E{"group", newUser.Group})
	}

	if len(newUser.Password) > 0 {
		setElements = append(setElements, bson.E{"password", user.HashPassword([]byte(newUser.Password))})
	}

	updatedUser := bson.M{"$set": setElements}

	updated, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updatedUser)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error updating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error updating user", 500))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	deleted, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error deleting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.NewError("Database error", "Error updating user", 500))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleted)
}
