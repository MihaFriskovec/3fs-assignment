package user

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/api/users", List).Methods("GET")
	// router.HandleFunc("/api/users", Create).Methods("POST")
	// router.HandleFunc("/api/users/{id}", Read).Methods("GET")
	// router.HandleFunc("/api/users/{id}", Update).Methods("PUT")
	// router.HandleFunc("/api/users/{id}", Delete).Methods("DELETE")

	return router
}
