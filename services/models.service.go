package services

import (
	"context"

	"github.com/Subasri-V/application-new/netxd_customer_dal/interfaces"
	models "github.com/Subasri-V/application-new/netxd_customer_dal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
	//	client *mongo.Client
}

func InitializeCustomerService(ctx context.Context, collection *mongo.Collection) interfaces.ICustomer {
	return &CustomerService{ctx, collection}
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

func (c*CustomerService) DeleteCustomerById(id int32) (*mongo.DeleteResult,error){
	filter := bson.D{{Key: "customerid", Value: id}}
	//var customer *models.Customer
	res, err := c.mongoCollection.DeleteOne(c.ctx, filter)
	
	if err != nil {
		return nil, err
	}
	return res, nil
}
