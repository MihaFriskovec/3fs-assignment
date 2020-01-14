package main

import (
	"log"
	"net/http"

	group "github.com/MihaFriskovec/3fs-assignment/groups/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", group.Routes()))
}
