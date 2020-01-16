package group

import (
	"testing"

	group "github.com/MihaFriskovec/3fs-assignment/groups/api"
	"github.com/MihaFriskovec/3fs-assignment/models"
)

var name = "Test name"

var demoGroup = models.Group{
	Name: &name,
}

func TestCreateGroup(t *testing.T) {
	err, inserted := group.Create(&demoGroup)

	if err != nil {
		t.Error("Error creating group")
	}

	if *inserted.Name != "Test name" {
		t.Error("Error saving Group.Name", *inserted.Name)
	}
}

func TestListGroup(t *testing.T) {
	err, list := group.List(1, 1, "", "")

	if err != nil {
		t.Error("Error listing groups")
	}

	if len(list) != 1 {
		t.Error("Group list should have one item")
	}
}

func TestReadGroup(t *testing.T) {
	err, inserted := group.Create(&demoGroup)

	err, item := group.Read(inserted.ID)

	if err != nil {
		t.Error("Error reading group")
	}

	if item.ID != inserted.ID {
		t.Error("Error reading group")
	}

	if *item.Name != "Test name" {
		t.Error("Error reading group")
	}
}

func TestUpdateGroup(t *testing.T) {
	err, inserted := group.Create(&demoGroup)

	updatedName := "Updated name"

	newGroup := models.Group{
		Name: &updatedName,
	}

	err, updated := group.Update(&newGroup, inserted.ID)

	if err != nil {
		t.Error("Error updating Group")
	}

	if updated.ModifiedCount != 1 {
		t.Error("No docuemnt updated")
	}

	err, updatedItem := group.Read(inserted.ID)

	if err != nil {
		t.Error("Error reading updated item", updatedItem)
	}

	if *updatedItem.Name != "Updated name" {
		t.Error("Error updating group name do not match", newGroup)
	}
}

func TestDeleteGroup(t *testing.T) {
	err, inserted := group.Create(&demoGroup)

	if err != nil {
		t.Error("Error", err)
	}

	err, deleted := group.Delete(inserted.ID)

	if err != nil {
		t.Error("Error", err)
	}

	if deleted.DeletedCount != 1 {
		t.Error("Error deleting group")
	}
}
