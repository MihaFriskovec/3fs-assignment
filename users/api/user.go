package user

import (
	"context"
	"errors"
	"log"

	"github.com/MihaFriskovec/3fs-assignment/db/helpers"
	"github.com/MihaFriskovec/3fs-assignment/models"
	user "github.com/MihaFriskovec/3fs-assignment/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func init() {
	userCollection = user.Users()
}

func Create(body *models.User) (error, *models.User) {
	newUser := user.User{}

	newUser.Name = body.Name
	newUser.Email = body.Email
	newUser.Password = user.HashPassword([]byte(*body.Password))
	newUser.Group = body.Group

	user, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Println("Error creating User", err)
		return errors.New("Error creating a new user"), nil
	}

	*body.Password = newUser.Password
	if oid, ok := user.InsertedID.(primitive.ObjectID); ok {
		body.ID = *&oid
	}

	return nil, body
}

func List(page int64, limit int64, sort string, project string) (error, []*models.User) {
	var usersList []*models.User

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

	cursor, err := userCollection.Find(context.TODO(), query, findOptions)

	if err != nil {
		log.Fatalln("Error listing Users", err)
		return errors.New("Error listing users"), nil
	}

	cursor.All(context.TODO(), &usersList)

	return nil, usersList
}

func Read(id primitive.ObjectID) (error, *models.User) {
	var user *models.User

	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		log.Println("Error reading User", err)
		return errors.New("Error reading User"), nil
	}

	return nil, user
}

func Update(body *models.User, id primitive.ObjectID) (error, *mongo.UpdateResult) {
	var setElements bson.D

	if len(*body.Name) > 0 {
		setElements = append(setElements, bson.E{"name", body.Name})
	}

	if len(*body.Email) > 0 {
		setElements = append(setElements, bson.E{"email", body.Email})
	}

	if len(body.Group) > 0 {
		setElements = append(setElements, bson.E{"user", body.Group})
	}

	if len(*body.Password) > 0 {
		password := user.HashPassword([]byte(*body.Password))
		setElements = append(setElements, bson.E{"password", password})
		*body.Password = password
	}

	updatedUser := bson.M{"$set": setElements}

	updated, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updatedUser)

	if err != nil {
		log.Println("Error updating User", err)
		return errors.New("Error updating User"), nil
	}

	if (updated.MatchedCount) == 0 {
		log.Println("Error updating User. No records found to update")
	}

	return nil, updated
}

func Delete(id primitive.ObjectID) (error, *mongo.DeleteResult) {
	deleted, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Println("Error deleting User", err)
		return errors.New("Error deleting User"), nil
	}

	return nil, deleted
}
