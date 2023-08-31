package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Subasri-V/application-new/netxd_customer_controller/constants"
	"github.com/Subasri-V/application-new/netxd_customer_dal/interfaces"
	"github.com/Subasri-V/application-new/netxd_customer_dal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
	client          *mongo.Client
}

func InitializeCustomerService(ctx context.Context, collection *mongo.Collection, client *mongo.Client) interfaces.ICustomer {
	return &CustomerService{ctx, collection, client}
}

func (c *CustomerService) CreateCustomer(customer *models.CustomerDetails) (*models.DBResponse, error) {
	res, err := c.mongoCollection.InsertOne(c.ctx, &customer)

	if err != nil {
		return nil, err
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.mongoCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (c *CustomerService) GetCustomerById(id int32) (*models.CustomerDetails, error) {
	filter := bson.M{"customerid": id}
	var customer *models.CustomerDetails
	err := c.mongoCollection.FindOne(c.ctx, filter).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerService) DeleteCustomerById(id int32) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "customerid", Value: id}}
	//var customer *models.Customer
	res, err := c.mongoCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CustomerService) UpdateCustomerById(id int32, customer *models.CustomerDetails) (*mongo.UpdateResult, error) {
	iv := bson.M{"customerid": id}
	fv := bson.M{"$set": &customer}
	res, err := c.mongoCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CustomerService) Transfer(id1 int32, id2 int32, amount int32) (string, error) {
	session, err := c.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(c.ctx)
	fmt.Println("ss")
	_, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {
		cust := c.client.Database(constants.DatabaseName).Collection("customer")
		trans := c.client.Database(constants.DatabaseName).Collection("transaction")
		fmt.Println("ss")

		res, err := cust.UpdateOne(ctx, bson.M{"customerid": id1}, bson.M{"$inc": bson.M{"balance": -amount}})
		if err != nil {
			return nil, err
		}
		fmt.Println("ss")

		fmt.Println(res)
		res, err = cust.UpdateOne(ctx, bson.M{"customerid": id2}, bson.M{"$inc": bson.M{"balance": amount}})
		if err != nil {
			return nil, err
		}
		res1, err := trans.InsertOne(ctx, &models.Transaction{Id: primitive.NewObjectID(), From: id1, To: id2, Amount: amount, TimeStamp: time.Now()})
		if err != nil {
			return nil, err
		}
		fmt.Println(res1)
		return res, nil

	})

	if err != nil {
		return "error occured", err
	}
	return "Transaction Successfull", nil
}
