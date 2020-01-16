package user

import "testing"

import user "github.com/MihaFriskovec/3fs-assignment/users/models"

func TestHashPassword(t *testing.T) {
	res := user.HashPassword([]byte("1234"))

	if res == "" {
		t.Error("Error hasing password")
	}
}
