package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserKey struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key     string             `json:"key,omitempty" bson:"key,omitempty"`
	KeyData primitive.ObjectID `json:"keyData,omitempty" bson:"keyData,omitempty"`
	Date    int64            `json:"date,omitempty" bson:"date,omitempty"`
}

type UserData struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Data      []DataIn           `json:"data,omitempty" bson:"data,omitempty"`
	Key       string             `json:"key,omitempty" bson:"key,omitempty"`
	LoginData int64              `json:"loginData,omitempty" bson:"loginData,omitempty"`
}

type DataIn struct {
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Store       string `json:"store,omitempty" bson:"store,omitempty"`
	Module      string `json:"module,omitempty" bson:"module,omitempty"`
	CurrentDate string `json:"currentDate,omitempty" bson:"currentDate,omitempty"`
	KW          string `json:"kw,omitempty" bson:"kw,omitempty"`
	Version     string `json:"version,omitempty" bson:"version,omitempty"`
	ID          int64  `json:"id,omitempty" bson:"id,omitempty"`
}

type MasterKey struct {
	Uk       UserKey                `json:"uk,omitempty" bson:"uk,omitempty"`
	Ud       UserData               `json:"ud,omitempty" bson:"ud,omitempty"`
	FirstId  *mongo.InsertOneResult `json:"firstId,omitempty" bson:"firstId,omitempty"`
	SecondId *mongo.InsertOneResult `json:"secondId,omitempty" bson:"secondId,omitempty"`
}
