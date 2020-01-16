package group

import (
	"testing"

	group "github.com/MihaFriskovec/3fs-assignment/groups/api"
	"github.com/MihaFriskovec/3fs-assignment/models"
)

func TestCreateGroup(t *testing.T) {
	var testGroup models.Group

	var name string = "Test name"

	testGroup.Name = &name

	err, inserted := group.Create(&testGroup)

	if err != nil {
		t.Error("Error creating group")
	}

	if *inserted.Name != name {
		t.Error("Error saving Group.Name")
	}
}
