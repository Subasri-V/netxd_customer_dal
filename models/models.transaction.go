package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id        primitive.ObjectID `json:"id" bson:"id"`
	From      int32              `json:"from" bson:"from"`
	To        int32              `json:"to" bson:"to"`
	Amount    int32              `json:"amount" bson:"amount"`
	TimeStamp time.Time          `json:"timeStamp" bson:"timeStamp"`
}
