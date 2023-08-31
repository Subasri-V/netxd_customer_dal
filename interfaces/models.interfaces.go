package interfaces

import (
	models "github.com/Subasri-V/application-new/netxd_customer_dal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICustomer interface {
	CreateCustomer(*models.CustomerDetails) (*models.DBResponse, error)
	GetCustomerById(id int32) (*models.CustomerDetails, error)
	DeleteCustomerById(id int32) (*mongo.DeleteResult, error)
}
