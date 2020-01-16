package helpers

import "testing"

import "github.com/MihaFriskovec/3fs-assignment/db/helpers"

func TestSort_ASC(t *testing.T) {
	res := helpers.Sort("name")

	if res == nil {
		t.Error("Error building sort options")
	}

	if res["name"] != 1 {
		t.Error("Error building sort options, name should be 1")
	}
}

func TestSort_DESC(t *testing.T) {
	res := helpers.Sort("-name")

	if res == nil {
		t.Error("Error building sort options")
	}

	if res["name"] != -1 {
		t.Error("Error building sort options, name should be -1")
	}
}

func TestProject(t *testing.T) {
	res := helpers.Project("name")

	if res == nil {
		t.Error("Error building project options")
	}

	if res["name"] != 1 {
		t.Error("Error building project options, name should be 1")
	}
}

func TestProjectMultiple(t *testing.T) {
	res := helpers.Project("name,_id")

	if res == nil {
		t.Error("Error building project options")
	}

	if res["name"] != 1 {
		t.Error("Error building project options, name should be 1")
	}

	if res["_id"] != 1 {
		t.Error("Error building project options, _id should be 1")
	}
}
