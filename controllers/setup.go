package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/foods", GetAllFood).Methods("GET")
	router.HandleFunc("/food/{id}", GetFood).Methods("GET")
	router.HandleFunc("/food", CreateFood).Methods("POST")
	router.HandleFunc("/food/{id}", UpdateFood).Methods("PUT")
	router.HandleFunc("/food/{id}", DeleteFood).Methods("DELETE")

	return router
}
