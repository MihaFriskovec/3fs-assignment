package group

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/api/groups", List).Methods("GET")
	// router.HandleFunc("/api/groups", Create).Methods("POST")
	// router.HandleFunc("/api/groups/{id}", Read).Methods("GET")
	// router.HandleFunc("/api/groups/{id}", Update).Methods("PUT")
	// router.HandleFunc("/api/groups/{id}", Delete).Methods("DELETE")

	return router
}
