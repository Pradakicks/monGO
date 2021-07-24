package getUsersV2

import (
	"context"
	"encoding/json"
	"fmt"
	"monGO-vibrisDB/helper"
	"monGO-vibrisDB/types"
	"net/http"

	"github.com/bradhe/stopwatch"
	"go.mongodb.org/mongo-driver/bson"
)

type APIv2 struct {}

func getUsers(w http.ResponseWriter, r *http.Request) {
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
		cur, err := collection.Find(context.TODO(), bson.M{})

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
		globalKeys = users
		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - UnAuthorized!"))
	}
}