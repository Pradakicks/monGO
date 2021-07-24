package main

import (
	"fmt"
	"log"
	helper "monGO-vibrisDB/helper"
	V1 "monGO-vibrisDB/v1"
	V2 "monGO-vibrisDB/v2"
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

	var v1 V1.APIv1 = V1.APIv1{Version: "v1"}

	var v2 V2.APIv2 = V2.APIv2{Version: "v2"}

	go helper.InitDB(&v1)
	helper.InitDB(&v2)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", v1.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.AddData).Methods("PATCH")
	router.HandleFunc("/api/v1/users", v1.AddUserKey).Methods("POST")

	router.HandleFunc("/api/v2/users", v2.GetUsers).Methods("GET")
	router.HandleFunc("/api/v2/users/{key}", v2.GetUser).Methods("GET")
	router.HandleFunc("/api/v2/users/{key}", v2.AddData).Methods("PATCH")
	router.HandleFunc("/api/v2/users", v2.AddUserKey).Methods("POST")

	err := http.ListenAndServe(":4123", router)
	log.Fatal(err)
}
