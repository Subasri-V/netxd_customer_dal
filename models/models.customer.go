package models

import (
	"time"
)

type CustomerDetails struct {
	Customerid int32     `json:"customerid" bson:"customerid"`
	Firstname  string    `json:"firstname" bson:"firstname"`
	Lastname   string    `json:"lastname" bson:"lastname"`
	Bankid     int32    `json:"bankid" bson:"bankid"`
	Balance    int32   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	IsActive   bool      `json:"isActive" bson:"isActive"`
}

type DBResponse struct {
	Customerid int32 `json:"customerid" bson:"customerid"`
	Firstname  string    `json:"firstname" bson:"firstname"`
	Balance    int32   `json:"balance" bson:"balance"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
    
}
