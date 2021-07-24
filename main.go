package main

import (
	"fmt"
	"log"
	helper "monGO-vibrisDB/helper"
	V1 "monGO-vibrisDB/v1"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabase interface {
	getVersion() string
	connectDB()
	connectDataDB()
	getCollection() *mongo.Collection
	getDataCollection() *mongo.Collection
}

type APIv2 struct{}

func main() {
	
	fmt.Println("Starting the application...")

	var v1 V1.APIv1 = V1.APIv1{
		Version: "v1",
	}

	helper.InitDB(&v1)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", v1.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.AddData).Methods("PATCH")
	router.HandleFunc("/api/v1/users", v1.AddUserKey).Methods("POST")

	err := http.ListenAndServe(":4123", router)
	log.Fatal(err)
}
