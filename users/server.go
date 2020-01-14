package main

import (
	"log"
	"net/http"

	user "github.com/MihaFriskovec/3fs-assignment/users/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", user.Routes()))
}
