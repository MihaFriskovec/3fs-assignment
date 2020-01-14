package helpers

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func Project(project string) bson.M {
	// Projection
	projectFields := strings.Split(project, ",")

	projectionQuery := bson.M{}

	for _, f := range projectFields {
		projectionQuery[f] = 1
	}

	return projectionQuery
}

func Sort(sort string) bson.M {
	var sortFlied string

	if len(sort) == 0 {
		return bson.M{}
	}

	sortOrder := sort[0:1]
	var order int

	if sortOrder == "-" || sortOrder == "+" {
		sortFlied = sort[1:]
		if sortOrder == "-" {
			order = -1
		}
	} else {
		order = 1
		sortFlied = sort
	}

	return bson.M{sortFlied: order}
}
