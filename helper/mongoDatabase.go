package helper

import (
	"context"
	"fmt"
	Interfaces "monGO-vibrisDB/interfaces"
	"monGO-vibrisDB/types"

	"github.com/bradhe/stopwatch"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllKeys(api Interfaces.MongoDatabase) []types.UserKey {
	watch := stopwatch.Start()

	var users []types.UserKey

	defer func() {
		watch.Stop()
		fmt.Printf("Length : %d Fetching Global Keys Took : %v\n", len(users), watch.Milliseconds())
	}()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := api.GetCollection().Find(context.TODO(), bson.M{})

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

func InitDB(api Interfaces.MongoDatabase) {
	fmt.Println("Initializing Version", api.GetVersion())
	go api.ConnectDB()
	go api.ConnectDataDB()
}

func GetKeyInPool(key string, api Interfaces.MongoDatabase) (bool, int) {
	var isPresent bool = false
	var position int = 999999
	for pos, value := range GetAllKeys(api) {
		if key == value.Key {
			isPresent = true
			position = pos
		}
	}
	return isPresent, position
}
