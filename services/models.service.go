package services

import (
	"Application-New/netxd_customer_dal/interfaces"
	"Application-New/netxd_customer_dal/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct{
	ctx context.Context
	mongoCollection *mongo.Collection
//	client *mongo.Client
}

func InitializeCustomerService(ctx context.Context,collection *mongo.Collection) interfaces.ICustomer {
	return &CustomerService{ctx,collection}
}

func (c*CustomerService)CreateCustomer(*models.CustomerDetails)(string){
	return "Successful"
}

