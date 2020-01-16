package group

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/MihaFriskovec/3fs-assignment/db/helpers"
	group "github.com/MihaFriskovec/3fs-assignment/groups/models"
	"github.com/MihaFriskovec/3fs-assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var groupCollection *mongo.Collection

func init() {
	groupCollection = group.Groups()
}

func Create(body *models.Group) (error, *models.Group) {
	newGroup := group.Group{}

	newGroup.Name = body.Name

	group, err := groupCollection.InsertOne(context.TODO(), newGroup)

	if err != nil {
		log.Println("Error creating Group", err)
		return errors.New("Error creating a new group"), nil
	}

	if oid, ok := group.InsertedID.(primitive.ObjectID); ok {
		body.ID = *&oid
	}

	return nil, body
}

func List(page int64, limit int64, sort string, project string) (error, []*models.Group) {
	var groupsList []*models.Group

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

	cursor, err := groupCollection.Find(context.TODO(), query, findOptions)

	if err != nil {
		log.Fatalln("Error listing Groups", err)
		return errors.New("Error listing groups"), nil
	}

	cursor.All(context.TODO(), &groupsList)
	fmt.Println(groupsList)

	return nil, groupsList
}

func Read(id primitive.ObjectID) (error, *models.Group) {
	var group *models.Group

	err := groupCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)

	if err != nil {
		log.Println("Error reading Group", err)
		return errors.New("Error reading Group"), nil
	}

	return nil, group
}

func Update(body *models.Group, id primitive.ObjectID) (error, *mongo.UpdateResult) {
	var newGroup group.Group

	var setElements bson.D

	if len(*body.Name) > 0 {
		setElements = append(setElements, bson.E{"name", newGroup.Name})
	}

	updatedGroup := bson.M{"$set": setElements}

	updated, err := groupCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updatedGroup)

	if err != nil {
		log.Println("Error updating Group", err)
		return errors.New("Error updating Group"), nil
	}

	return nil, updated
}

func Delete(id primitive.ObjectID) (error, *mongo.DeleteResult) {
	deleted, err := groupCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Println("Error deleting Group", err)
		return errors.New("Error deleting Group"), nil
	}

	return nil, deleted
}
