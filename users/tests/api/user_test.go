package user

import (
	"testing"

	"github.com/MihaFriskovec/3fs-assignment/models"
	user "github.com/MihaFriskovec/3fs-assignment/users/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var name = "Test name"
var pass = "123456"
var email = "test@test.si"
var group = primitive.NewObjectID()

var demoUser = models.User{
	Name:     &name,
	Password: &pass,
	Email:    &email,
	Group:    group,
}

func TestCreateGroup(t *testing.T) {
	err, inserted := user.Create(&demoUser)

	if err != nil {
		t.Error("Error creating user")
	}

	if *inserted.Name != name {
		t.Error("Error saving User.Name", *inserted.Name)
	}

	if *inserted.Name == pass {
		t.Error("Error hasih User.Password", *inserted.Password)
	}

	if inserted.Password == nil {
		t.Error("Error saving User.Password", *inserted.Password)
	}

	if *inserted.Email != email {
		t.Error("Error saving User.Email", *inserted.Email)
	}

	if inserted.Group != group {
		t.Error("Error saving User.Group", inserted.Group)
	}
}

func TestListUser(t *testing.T) {
	err, list := user.List(1, 1, "", "")

	if err != nil {
		t.Error("Error listing users")
	}

	if len(list) != 1 {
		t.Error("User list should have one item")
	}
}

func TestReadUser(t *testing.T) {
	err, inserted := user.Create(&demoUser)

	err, item := user.Read(inserted.ID)

	if err != nil {
		t.Error("Error reading user")
	}

	if item.ID != inserted.ID {
		t.Error("Error reading user")
	}

	if *item.Name != name {
		t.Error("Error reading user")
	}
}

func TestUpdateUser(t *testing.T) {
	err, inserted := user.Create(&demoUser)

	updatedName := "Updated name"
	updatedPass := "654321"
	updatedEmail := "test1@test.si"

	newUser := models.User{
		Name:     &updatedName,
		Password: &updatedPass,
		Email:    &updatedEmail,
	}

	err, updated := user.Update(&newUser, inserted.ID)

	if err != nil {
		t.Error("Error updating user")
	}

	if updated.ModifiedCount != 1 {
		t.Error("No docuemnt updated")
	}

	err, updatedItem := user.Read(inserted.ID)

	if err != nil {
		t.Error("Error reading updated item", updatedItem)
	}

	if *updatedItem.Name != updatedName {
		t.Error("Error updating group name do not match", *updatedItem.Name)
	}

	if *updatedItem.Password != updatedPass {
		t.Error("Error updating group password do not match", *updatedItem.Password)
	}

	if *updatedItem.Email != updatedEmail {
		t.Error("Error updating group email do not match", *updatedItem.Email)
	}

}

func TestDeleteUser(t *testing.T) {
	err, inserted := user.Create(&demoUser)

	if err != nil {
		t.Error("Error", err)
	}

	err, deleted := user.Delete(inserted.ID)

	if err != nil {
		t.Error("Error", err)
	}

	if deleted.DeletedCount != 1 {
		t.Error("Error deleting user")
	}
}
