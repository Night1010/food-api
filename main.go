package main

import (
	"net/http"
	"project/controllers"
	"project/models"
)

func main() {

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()

}
