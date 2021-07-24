package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	helper "monGO-vibrisDB/helper"
	"monGO-vibrisDB/types"
	"net/http"
	"time"

	"github.com/bradhe/stopwatch"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase interface {
	getVersion() string
	connectDB()
	connectDataDB()
	getCollection() *mongo.Collection
	getDataCollection() *mongo.Collection
}

type APIv2 struct {}

type APIv1 struct {
	version string
	collection *mongo.Collection
	dataCollection *mongo.Collection
}

func (a *APIv1) getVersion() string {
	return a.version
}
func (v *APIv1) connectDB () {
	v.collection = helper.ConnectDB()
}

func (v *APIv1) connectDataDB () {
	v.dataCollection = helper.ConnectDBData()
}

func (v *APIv1) getCollection () *mongo.Collection {
	return v.collection
}

func (v *APIv1) getDataCollection () *mongo.Collection {
	return v.dataCollection
}

func (a *APIv1) getUsers(w http.ResponseWriter, r *http.Request) {

	watch := stopwatch.Start()

	defer func() {
		watch.Stop()
		fmt.Printf("Request Took : %v\n", watch.Milliseconds())
	}()
	val, ok := r.Header["Apiekdfudks9"]
	if ok && val[0] == "secretVibNoa9o73jd91kd0akd8nf38ald8nfoa8dnalkjsd98fkksd8fnalsdfha9sdfnasdp;fpasdjhfpioashdf9asdhfasdlfasd8fasdofbasdkjf" {
		fmt.Println("Validated")
		var users []types.UserKey

		// bson.M{},  we passed empty filter. So we want to get all data.
		cur, err := a.collection.Find(context.TODO(), bson.M{})

		if err != nil {
			helper.GetError(err, w)
			return
		}

		// Close the cursor once finished
		/*A defer statement defers the execution of a function until the surrounding function returns.
		simply, run cur.Close() process but after cur.Next() finished.*/
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var user types.UserKey
			// & character returns the memory address of the following variable.
			err := cur.Decode(&user) // decode similar to deserialize process.
			if err != nil {
				fmt.Println(err)
			}

			// add item our array
			users = append(users, user)
		}

		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - UnAuthorized!"))
	}
}

func (a *APIv1) addUserKey(w http.ResponseWriter, r *http.Request) {
	watch := stopwatch.Start()

	w.Header().Add("content-type", "application/json")
	var userkey types.UserKey
	json.NewDecoder(r.Body).Decode(&userkey)
	userkey.Date = time.Now().UTC().UnixNano()
	t := time.Now()
	currentFormatDate := t.Format("2006-01-02")
	datain := []types.DataIn{
		{
			Type:        "TEST",
			Store:       "TEST",
			KW:          "TEST",
			CurrentDate: currentFormatDate,
			Module:      "TEST",
			Version:     "TEST",
			ID:          userkey.Date,
		},
	}
	var userdata types.UserData = types.UserData{
		ID:        primitive.NewObjectID(),
		Key:       userkey.Key,
		LoginData: time.Now().UTC().UnixNano(),
		Data:      datain,
	}
	defer func() {
		watch.Stop()
		fmt.Printf("Adding User Key : %s Request Took : %v\n", userkey.Key, watch.Milliseconds())
	}()
	userkey.KeyData = userdata.ID
	isPresent, _ := getKeyInPool(userkey.Key, a)
	if !isPresent {
		result, err := a.collection.InsertOne(context.TODO(), userkey)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		res, err := a.dataCollection.InsertOne(context.TODO(), userdata)
		if err != nil {
			helper.GetError(err, w)
			return
		}

		var masterResult types.MasterKey = types.MasterKey{
			Uk:       userkey,
			Ud:       userdata,
			FirstId:  result,
			SecondId: res,
		}
		json.NewEncoder(w).Encode(masterResult)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Key Already Present!"))
	}

}
func (a *APIv1) getUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	watch := stopwatch.Start()

	defer func() {
		watch.Stop()
		fmt.Printf("Request Took : %v\n", watch.Milliseconds())
	}()

	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	key := params["key"]
	val, ok := r.Header["Apiekdfudks9"]

	if ok && val[0] == "secretVibNoa9o73jd91kd0akd8nf38ald8nfoa8dnalkjsd98fkksd8fnalsdfha9sdfnasdp;fpasdjhfpioashdf9asdhfasdlfasd8fasdofbasdkjf" {
		fmt.Println(key, "Validated")
		isPresent, pos := getKeyInPool(key, a)
		if !isPresent {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Key Not Found!"))
		} else {
			json.NewEncoder(w).Encode(getAllKeys(a)[pos])
		}

		// var user types.UserKey
		// we get params with mux.

		// string to primitive.ObjectID

		// // We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
		// filter := bson.M{"key": key}
		// err := collection.FindOne(context.TODO(), filter).Decode(&user)

		// 	helper.GetError(err, w)
		// 	return
		// }
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - UnAuthorized!"))
	}
}

func (a *APIv1) addData(w http.ResponseWriter, r *http.Request) {
	watch := stopwatch.Start()

	
	w.Header().Add("content-type", "application/json")
	var params = mux.Vars(r)

	key := params["key"]
	defer func() {
		watch.Stop()
		fmt.Printf("Adding Data to %s: Request Took : %v\n", key, watch.Milliseconds())
	}()
	isPresent, _ := getKeyInPool(key, a)
	if !isPresent {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Key Not Found!"))
	} else {
		var currentData types.DataIn
		opts := options.FindOneAndUpdate().SetUpsert(true)
		json.NewDecoder(r.Body).Decode(&currentData)
		filter := bson.M{"key": key}
		update := bson.M{
			"$push": bson.M{"data": currentData},
		}
		var updatedDocument bson.M
		err := a.dataCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(updatedDocument)

	}
}

func getAllKeys(api mongoDatabase) []types.UserKey {
	watch := stopwatch.Start()

	var users []types.UserKey

	defer func() {
		watch.Stop()
		fmt.Printf("Length : %d Fetching Global Keys Took : %v\n", len(users), watch.Milliseconds())
	}()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := api.getCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {

		var user types.UserKey

		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			fmt.Println(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	return users
}

func initDB(api mongoDatabase) {
	fmt.Println("Initializing Version", api.getVersion())
	go api.connectDB()
	go api.connectDataDB()
}

func getKeyInPool(key string, api mongoDatabase) (bool, int) {
	var isPresent bool = false
	var position int = 999999
	for pos, value := range getAllKeys(api) {
		if key == value.Key {
			isPresent = true
			position = pos
		}
	}
	return isPresent, position
}

func main() {
	fmt.Println("Starting the application...")

	var v1 APIv1 = APIv1{
		version: "v1",
	}

	initDB(&v1)

	// mongodb+srv://vibris-User:eIDpR4kttFu57FHE@vibris.jyxhh.mongodb.net/VibrisData?retryWrites=true&w=majority

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", v1.getUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.getUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{key}", v1.addData).Methods("PATCH")
	router.HandleFunc("/api/v1/users", v1.addUserKey).Methods("POST")


	err := http.ListenAndServe(":4123", router)
	log.Fatal(err)
}
